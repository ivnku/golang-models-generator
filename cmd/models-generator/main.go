package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"models-generator/internal/adapters"
	"models-generator/internal/adapters/sqlserver"
	"models-generator/internal/initApp"
	_ "os"
)

func main() {
	flags := initApp.InitFlags()
	config, err := initApp.InitConfig(*flags.ConfigPath)

	fmt.Printf("Config is %v \n", config)

	if err != nil {
		log.Fatal(err)
	}

	var adapter adapters.IAdapter
	adapter = &sqlserver.SqlServerAdapter{}
	db, err := adapter.GetDB(config)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var res []adapters.AdapterResultSet
	err = db.Select(&res, adapter.GetSql("Backend_Anwendung"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
}
