package engine

import (
	"fmt"
	"log"
	"models-generator/config"
	"models-generator/internal/adapters"
	"models-generator/internal/adapters/base"
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
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
	return nil
}
