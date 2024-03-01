package main

import (
	"fmt"
	"reflect"
)

const (
	ID       = 1 << iota // 1 << 0 == 1
	Username             // 1 << 1 == 2
	Email                // 1 << 2 == 4
	Address              // 1 << 3 == 8
	Status               // 1 << 4 == 16
)

type User struct {
	ID       int    `db:"id" db_ops:"create"`
	Username string `db:"username" db_ops:"create,update"`
	Email    string `db:"email" db_ops:"create,update"`
	Address  string `db:"address" db_ops:"update"`
	Status   int    `db:"status" db_ops:"create,update"`
	Delete   string `db:"delete" db_ops:"delete"`
}

func GetFieldsPointers(u interface{}, args ...func(*[]reflect.StructField)) []interface{} {
	v := reflect.ValueOf(u).Elem()
	t := v.Type()
	res := []interface{}{}
	var sl []reflect.StructField
	for i := 0; i < v.NumField(); i++ {
		fF := t.Field(i)
		sl = append(sl, fF)
	}
	for _, arg := range args {
		arg(&sl)
	}
	for i := 0; i < len(sl); i++ {
		tag := sl[i].Tag.Get("db")
		for j := 0; j < v.NumField(); j++ {
			fieldT := v.Type().Field(j)
			tagG := fieldT.Tag.Get("db")
			if tagG == tag {
				f := v.Field(j).Addr().Interface()
				res = append(res, f)
			}
		}
	}

	return res
}

func main() {
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	filter1 := func(fields *[]reflect.StructField) {
		var res []reflect.StructField
		requiredFields := []int{ID, Username, Email}

		for i := range requiredFields {
			for j := range *fields {
				val := requiredFields[i]
				idx := 1 << j

				if val&idx != 0 {
					res = append(res, (*fields)[j])
					break
				}
			}
		}

		*fields = res
	}

	filter2 := func(fields *[]reflect.StructField) {
		var res []reflect.StructField

		for i := range *fields {
			if (*fields)[i].Tag.Get("db_ops") != "create" {
				res = append(res, (*fields)[i])
			}
		}

		*fields = res
	}

	pointers := GetFieldsPointers(&user, filter1, filter2)

	for _, pointer := range pointers {
		switch v := pointer.(type) {
		case *int:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		case *string:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		}
	}

	fmt.Println()

	pointers = GetFieldsPointers(&user, FilterByFields(ID, Username, Email))
	fmt.Println("FilterByFields(ID, Username, Email)")

	for _, pointer := range pointers {
		switch v := pointer.(type) {
		case *int:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		case *string:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		}
	}

	filterTag := map[string]func(value string) bool{
		"db": func(value string) bool {
			values := []string{"username", "address", "status"}

			for _, v := range values {
				if v == value {
					return true
				}
			}
			return false
		},
	}

	fmt.Println()

	pointers = GetFieldsPointers(&user, FilterByTags(filterTag))
	fmt.Println("FilterByTags(filterTag)")

	for _, pointer := range pointers {
		switch v := pointer.(type) {
		case *int:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		case *string:
			fmt.Printf("%T %v: %v\n", pointer, pointer, *v)
		}
	}
}

func FilterByFields(fields ...int) func(fieldss *[]reflect.StructField) {
	return func(fieldss *[]reflect.StructField) {
		var res []reflect.StructField

		for i := range *fieldss {
			for _, requiredField := range fields {
				if requiredField == i {
					res = append(res, (*fieldss)[i])
					break
				}
			}
		}

		*fieldss = res
	}
}

func FilterByTags(tags map[string]func(value string) bool) func(fields *[]reflect.StructField) {
	return func(fields *[]reflect.StructField) {
		var res []reflect.StructField

		for i := range *fields {
			tag := (*fields)[i].Tag.Get("db")
			for _, filterFunc := range tags {
				if filterFunc(tag) {
					res = append(res, (*fields)[i])
					break
				}
			}
		}

		*fields = res
	}
}
