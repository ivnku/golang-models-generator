package engine

import (
	"fmt"
	"models-generator/config"
	"models-generator/internal/adapters"
	"models-generator/internal/adapters/base"
	"models-generator/internal/engine/model"
	"models-generator/internal/initApp"
)

func Generate(config *config.AppConfig, flags *initApp.Flags) error {
	var adapter base.IAdapter
	adapter, err := adapters.GetAdapter(config.Connection.Driver)

	if err != nil {
		return fmt.Errorf("could't create adapter: %s", err.Error())
	}

	db, err := adapter.GetDB(config)
	defer db.Close()

	if err != nil {
		return fmt.Errorf("could't create db connection: %s", err.Error())
	}

	var res []base.AdapterResultSet
	err = db.Select(&res, adapter.GetSql(*flags.Table))

	if err != nil {
		return fmt.Errorf("couldn't get tables' data from database: %s", err.Error())
	}

	err = aggregateModelsData(res, flags, &adapter)

	if err != nil {
		return fmt.Errorf("couldn't aggregate models data: %s", err.Error())
	}

	fmt.Printf("%#v\n", res)
	return nil
}

func aggregateModelsData(data []base.AdapterResultSet, flags *initApp.Flags, adapter *base.IAdapter) error {
	packageName, err := model.GetPackageName(*flags.Path)

	if err != nil {
		return err
	}

	groupedData := groupResultSet(packageName, data, adapter)

	fmt.Printf("%v", groupedData)

	return nil
}

func groupResultSet(packageName string, data []base.AdapterResultSet, adapter *base.IAdapter) map[string]*model.ModelData {
	grouped := make(map[string]*model.ModelData)
	for _, row := range data {
		if _, ok := grouped[row.Table]; !ok {
			columns := []model.ColumnData{
				{
					Name: model.GetCamelCaseName(row.Column),
					Type: model.GetType(row, (*adapter).GetTypesMapping()),
					Tag:  model.GetTag(row),
				},
			}
			grouped[row.Table] = &model.ModelData{packageName, model.GetCamelCaseName(row.Table), columns}
		} else {
			column := model.ColumnData{
				Name: model.GetCamelCaseName(row.Column),
				Type: model.GetType(row, (*adapter).GetTypesMapping()),
				Tag:  model.GetTag(row),
			}
			entry := grouped[row.Table]
			entry.Columns = append(entry.Columns, column)
		}
	}
	return grouped
}
