package initApp

import "flag"

type Flags struct {
	Table        *string
	GeneratePath *string
	ConfigPath   *string
}

func InitFlags() *Flags {
	f := &Flags{
		Table: flag.String(
			"t",
			"",
			"the name of the table you need to create a model for. If no name specified - generate models for all tables",
		),
		GeneratePath: flag.String(
			"p",
			"",
			"the path where to generate models",
		),
		ConfigPath: flag.String(
			"c",
			"./config.yml",
			"the path to the config file, './config.yml' by default",
		),
	}

	flag.Parse()

	return f
}
