package engine

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"models-generator/config"
	"models-generator/internal/adapters"
	"models-generator/internal/adapters/base"
	"models-generator/internal/engine/model"
	"models-generator/internal/initApp"
	"os"
	"text/template"
)

func Generate(config *config.AppConfig, flags *initApp.Flags) error {
	var generatePath string
	if *flags.GeneratePath != "" {
		generatePath = *flags.GeneratePath
	} else if config.GeneratedModelsPath != "" {
		generatePath = config.GeneratedModelsPath
	} else {
		generatePath = "./models"
	}
	*flags.GeneratePath = generatePath

	groupedData, err := aggregateModelsData(flags, config)

	if err != nil {
		return fmt.Errorf("couldn't aggregate models data: %s", err.Error())
	}

	templString, err := os.ReadFile(config.TemplatePath)

	if err != nil {
		return fmt.Errorf("couldn't find a template file: %s", err.Error())
	}

	templ, err := template.New("model").Parse(string(templString))
	if err != nil {
		return fmt.Errorf("couldn't parse a template string: %s", err.Error())
	}

	//_ = math.Ceil(float64(len(groupedData)) / 10)
	for _, value := range groupedData {
		err = createModelFile(value, templ, flags)
		if err != nil {
			fmt.Printf("couldn't create model file: %s\n", err.Error())
		}
	}

	return nil
}

func aggregateModelsData(flags *initApp.Flags, config *config.AppConfig) (map[string]*model.ModelData, error) {
	packageName, err := model.GetPackageName(*flags.GeneratePath)

	if err != nil {
		return nil, err
	}

	var adapter base.IAdapter
	adapter, err = adapters.GetAdapter(config.Connection.Driver)

	if err != nil {
		return nil, fmt.Errorf("could't create adapter: %s", err.Error())
	}

	db, err := adapter.GetDB(config)
	defer db.Close()

	if err != nil {
		return nil, fmt.Errorf("could't create db connection: %s", err.Error())
	}

	var data []base.AdapterResultSet
	err = db.Select(&data, adapter.GetSql(*flags.Table))

	if err != nil {
		return nil, fmt.Errorf("couldn't get tables' data from database: %s", err.Error())
	}

	groupedData := groupResultSet(packageName, data, &adapter)

	return groupedData, nil
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
			modelData := &model.ModelData{columns, make(map[string]struct{}), packageName, model.GetCamelCaseName(row.Table)}
			model.AddImport(modelData, &columns[0])
			grouped[row.Table] = modelData
		} else {
			column := model.ColumnData{
				Name: model.GetCamelCaseName(row.Column),
				Type: model.GetType(row, (*adapter).GetTypesMapping()),
				Tag:  model.GetTag(row),
			}
			modelData := grouped[row.Table]
			model.AddImport(modelData, &column)
			grouped[row.Table].Columns = append(grouped[row.Table].Columns, column)
		}
	}
	return grouped
}

func createModelFile(model *model.ModelData, templ *template.Template, flags *initApp.Flags) error {
	//time.Sleep(500 * time.Millisecond)

	err := os.MkdirAll(*flags.GeneratePath, 0755)

	if err != nil {
		return fmt.Errorf("couldn't create a path for a model file for %s: %s", model.ModelName, err.Error())
	}

	modelFile, err := os.Create(fmt.Sprintf("%s/%s.go", *flags.GeneratePath, model.ModelName))
	defer modelFile.Close()

	if err != nil {
		return fmt.Errorf("couldn't create a model file for %s: %s", model.ModelName, err.Error())
	}

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	err = templ.Execute(buf, model)

	if err != nil {
		return fmt.Errorf("couldn't execute a template for a model file %s: %s", model.ModelName, err.Error())
	}

	w := bufio.NewWriter(modelFile)
	formatted, err := format.Source(buf.Bytes())

	if err != nil {
		return fmt.Errorf("couldn't format the model file %s: %s", model.ModelName, err.Error())
	}

	_, err = w.Write(formatted)

	if err != nil {
		return fmt.Errorf("couldn't write to the model file %s: %s", model.ModelName, err.Error())
	}

	err = w.Flush()

	if err != nil {
		return fmt.Errorf("couldn't perform Flush() on a model file %s: %s", model.ModelName, err.Error())
	}

	fmt.Printf("Model for %s created!\n", model.ModelName)
	return nil
}
