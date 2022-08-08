package sqlserver

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"models-generator/config"
)

type SqlServerAdapter struct{}

func (a *SqlServerAdapter) GetTypesMapping() map[string]string {
	return map[string]string{
		"int":       "int",
		"varchar":   "string",
		"bit":       "bool",
		"decimal":   "float32",
		"money":     "float32",
		"datetime":  "time.Time",
		"datetime2": "time.Time",
	}
}

func (a *SqlServerAdapter) GetSql(table string) string {
	var tableCond string

	if table != "" {
		tableCond = fmt.Sprintf("WHERE tab.name = '%s'", table)
	}

	return fmt.Sprintf(`
		SELECT
			tab.name, col.name AS [column], type.name as type, col.is_nullable, col.max_length, col.precision, col.scale, 
			CASE WHEN tab2.name IS NULL THEN 0 ELSE 1 END as is_foreign_key, i.is_primary_key, tab2.name AS referenced_table, 
			col2.name AS referenced_column
		FROM sys.tables tab
		INNER JOIN sys.columns col ON tab.object_id = col.object_id
		LEFT JOIN sys.foreign_key_columns foreignCol ON col.column_id = foreignCol.parent_column_id AND foreignCol.parent_object_id = tab.object_id
		LEFT JOIN sys.tables tab2 ON tab2.object_id = foreignCol.referenced_object_id
		LEFT JOIN sys.columns col2 ON col2.column_id = foreignCol.referenced_column_id AND col2.object_id = tab2.object_id
		LEFT JOIN sys.index_columns ic ON col.column_id = ic.column_id AND ic.object_id = tab.object_id
		LEFT JOIN sys.indexes i ON ic.index_id=i.index_id AND i.object_id = tab.object_id
		INNER JOIN sys.types type ON col.system_type_id = type.system_type_id
		%s
		ORDER BY tab.name
	`, tableCond)
}

func (a *SqlServerAdapter) GetDB(config *config.AppConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlserver", config.Connection.Dsn)
	if err != nil {
		err = fmt.Errorf("couldn't create db connection! Check your dsn string in the config. Error: %s", err.Error())
		return nil, err
	}
	return db, nil
}
