package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"models-generator/internal/engine"
	"models-generator/internal/initApp"
	_ "os"
)

func main() {
	flags := initApp.InitFlags()
	config, err := initApp.InitConfig(*flags.ConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	err = engine.Generate(config, flags)

	if err != nil {
		log.Fatal(err)
	}
}
