//go:build dev

package util

import (
	"fmt"
	"reflect"
	"strings"
)

func PrintStruct(v interface{}, prefix string) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Printf("Expected a struct, got %s\n", val.Kind())
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Name

		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		// Use strings.HasPrefix instead of StartsWith
		if field.Kind() == reflect.Struct && !strings.HasPrefix(field.Type().PkgPath(), "time") {
			PrintStruct(field.Interface(), fieldName)
		} else {
			fmt.Printf("%s: Type(%T) Value(%v)\n", fieldName, field.Interface(), field.Interface())
		}
	}
}
