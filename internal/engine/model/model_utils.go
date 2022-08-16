package model

import (
	"fmt"
	"github.com/oriser/regroup"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"models-generator/internal/adapters/base"
	"strings"
)

type ModelData struct {
	PackageName string
	ModelName   string
	Columns     []ColumnData
}

type ColumnData struct {
	Name string
	Type string
	Tag  string
}

func GetPackageName(path string) (string, error) {
	re := regroup.MustCompile("[\\\\\\/](?P<Package>[a-zA-Z-_0-9]+)$")
	match, err := re.Groups(path)

	if err != nil {
		return "", fmt.Errorf("couldn't get package name: %s", err.Error())
	}

	return match["Package"], nil
}

func GetCamelCaseName(value string) string {
	splitted := strings.Split(value, "_")
	var result string
	caser := cases.Title(language.English)
	for _, s := range splitted {
		result += caser.String(s)
	}
	return result
}

func GetType(columnData base.AdapterResultSet, typesMapping map[string]string) string {
	mappedType := typesMapping[columnData.Type]
	if columnData.IsPrimaryKey.Bool && mappedType == "int" {
		return "uint64"
	}
	if columnData.IsNullable {
		switch mappedType {
		case "int":
			return "sql.NullInt32"
		case "string":
			return "sql.NullString"
		case "bool":
			return "sql.NullBool"
		case "float32":
			return "sql.NullFloat64"
		case "time.Time":
			return "sql.NullTime"
		default:
			return "sql.NullString"
		}
	} else {
		return mappedType
	}
}

func GetTag(columnData base.AdapterResultSet) string {
	return fmt.Sprintf("`db:\"%s\"`", columnData.Column)
}
