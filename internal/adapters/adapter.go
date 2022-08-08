package adapters

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"models-generator/config"
)

type IAdapter interface {
	GetTypesMapping() map[string]string
	// GetSql returns raw select sql query and should have the following columns in a result set:
	// 'name', 'column', 'type', 'is_nullable', 'max_length', 'precision', 'scale',
	// 'is_foreign_key', 'is_primary_key', 'referenced_table', 'referenced_column'
	GetSql(table string) string
	GetDB(config *config.AppConfig) (*sqlx.DB, error)
}

type AdapterResultSet struct {
	Name             string
	Column           string
	Type             string
	ReferencedTable  sql.NullString `db:"referenced_table"`
	ReferencedColumn sql.NullString `db:"referenced_column"`
	MaxLength        sql.NullInt32  `db:"max_length"`
	Precision        sql.NullInt32
	Scale            sql.NullInt32
	IsPrimaryKey     sql.NullBool `db:"is_primary_key"`
	IsNullable       bool         `db:"is_nullable"`
	IsForeignKey     bool         `db:"is_foreign_key"`
}
