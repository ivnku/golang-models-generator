package base

import (
	"database/sql"
	"fmt"
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

type Adapter struct{}

func (a *Adapter) GetDB(config *config.AppConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(config.Connection.Driver, config.Connection.Dsn)
	if err != nil {
		err = fmt.Errorf("couldn't create db connection! Check your dsn string in the config. Error: %s", err.Error())
		return nil, err
	}
	return db, nil
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
