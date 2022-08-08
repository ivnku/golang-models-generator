package initApp

import "flag"

type Flags struct {
	Table      *string
	Path       *string
	ConfigPath *string
}

func InitFlags() *Flags {
	f := &Flags{
		Table: flag.String(
			"t",
			"none",
			"the name of the table you need to create a model for. If no name specified - generate models for all tables",
		),
		Path: flag.String(
			"p",
			"./models",
			"the path where to generate models",
		),
		ConfigPath: flag.String(
			"c",
			"./config.yml",
			"the path to config file, './config.yml' by default",
		),
	}

	flag.Parse()

	return f
}
