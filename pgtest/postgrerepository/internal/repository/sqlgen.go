package repository

import (
	"fmt"
	"reflect"
	"strings"
)

type Tabler interface {
	TableName() string
}

type StructInfo struct {
	Fields   []string
	Pointers []interface{}
}

type SQLgen interface {
	CreateTableSQL(table Tabler) string
}

type SQLiteGen struct{}

func (sqlgen *SQLiteGen) CreateTableSQL(table Tabler) string {
	quiery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", table.TableName())
	fields := getFields(table)
	for i, field := range fields {
		quiery += fmt.Sprintf(" %s %s,\n", GetstructInfo(table).Fields[i], field.Tag.Get("db_type"))
	}

	quiery = strings.TrimSuffix(quiery, ",\n") + "\n"
	return quiery
}

func getFields(v interface{}) []reflect.StructField {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fields := make([]reflect.StructField, 0)

	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i))

	}

	return fields

}

func GetstructInfo(u interface{}, args ...func(*[]reflect.StructField)) StructInfo {
	value := reflect.ValueOf(u).Elem()
	structFields := make([]reflect.StructField, 0)

	for i := 0; i < value.NumField(); i++ {
		structFields = append(structFields, value.Type().Field(i))
	}

	for i := range args {
		if args[i] != nil {
			args[i](&structFields)
		}
	}

	strInf := StructInfo{}

	for _, field := range structFields {
		vl := value.FieldByName(field.Name)
		strInf.Fields = append(strInf.Fields, field.Tag.Get("db"))
		strInf.Pointers = append(strInf.Pointers, vl.Addr().Interface())

	}

	return strInf
}
