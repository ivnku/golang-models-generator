package main

import (
	"flag"
	"fmt"
	_ "os"
)

func main() {
	tableName := flag.String(
		"t",
		"none",
		"a name of the table you need to create a model for. If no name specified - generate models for all tables",
	)
	flag.Parse()
	fmt.Printf("flags are: %v", *tableName)
}
