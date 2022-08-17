package adapters

import (
	"fmt"
	"models-generator/internal/adapters/base"
	"models-generator/internal/adapters/sqlserver"
)

func GetAdapter(driver string) (base.IAdapter, error) {
	switch driver {
	case "sqlserver":
		return &sqlserver.SqlServerAdapter{}, nil
	default:
		return nil, fmt.Errorf("there is no such driver as '%s'", driver)
	}
}
