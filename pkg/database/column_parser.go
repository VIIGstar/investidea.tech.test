package database

import (
	"fmt"
	"gorm.io/gorm/schema"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var DateTimeColumn = map[string]bool{
	"created_at": true,
	"updated_at": true,
	"deleted_at": true,
}

func SelectColumns(tabler schema.Tabler, options ...ColumnParseOption) string {
	var (
		ignoreFields map[string]bool
	)
	for _, opt := range options {
		switch impl := opt.(type) {
		case IgnoreColumnsOption:
			ignoreFields = getMapIgnoreFields(impl)
		}
	}

	listColumns := getAllFields(tabler)
	prefixLIstColumn := make([]string, 0)
	for _, col := range listColumns {
		if ignoreFields[col] {
			continue
		}
		prefixLIstColumn = append(prefixLIstColumn, fmt.Sprintf("%v.%v", tabler.TableName(), col))
	}

	return strings.Join(prefixLIstColumn, ",")
}

func getAllFields(obj interface{}) []string {
	listColumns := []string{}
	typeObj := reflect.TypeOf(obj)
	valueObj := reflect.ValueOf(obj)

	for i := 0; i < typeObj.NumField(); i++ {
		fieldName := ToSnakeCase(typeObj.Field(i).Name)
		switch typeObj.Field(i).Type.Kind() {
		case reflect.Struct, reflect.Ptr:
			if DateTimeColumn[fieldName] {
				listColumns = append(listColumns, fieldName)
				continue
			}
			listColumns = append(listColumns, getAllFields(valueObj.Field(i).Interface())...)
			continue
		}

		tagValue := typeObj.Field(i).Tag.Get("gorm")
		if tagValue == "-" {
			continue
		}
		// check column defined
		idx := strings.Index(tagValue, "column:")
		if idx < 0 {
			listColumns = append(listColumns, fieldName)
			continue
		}

		colStart := tagValue[idx+len("column:"):]
		parts := strings.Split(colStart, ";")
		listColumns = append(listColumns, parts[0])
	}

	return listColumns
}

func getMapIgnoreFields(option IgnoreColumnsOption) map[string]bool {
	ignoreMap := map[string]bool{}
	for _, field := range option.Fields {
		ignoreMap[ToSnakeCase(field)] = true
	}

	return ignoreMap
}
