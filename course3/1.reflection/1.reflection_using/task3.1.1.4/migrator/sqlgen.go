package migrator

import (
	"fmt"
	"reflect"
	"strings"
	"studentgit.kata.academy/fcons/go-kata/course3/1.reflection/1.reflection_using/task3.1.1.4/tabler"
)

type SQLGenerator interface {
	CreateTableSQL(table tabler.Tabler) string
}

type SQLiteGenerator struct{}

func (sg *SQLiteGenerator) CreateTableSQL(table tabler.Tabler) string {
	ref := reflect.TypeOf(table).Elem()
	res := ""
	for i := 0; i < ref.NumField(); i++ {
		res += fmt.Sprintf("%s %s, ", ref.Field(i).Tag.Get("db"), ref.Field(i).Tag.Get("db_type"))
	}
	res = strings.TrimSuffix(res, ", ")
	cts := fmt.Sprintf("CREATE TABLE IF NOT EXISTS '%s' ('%s')", table.TableName(), res)
	return cts
}
