package helper

import (
	"reflect"
)

func JsonResponse(status bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{} {"status": status, "message": message, "data": data}
}

func DynamicWhere() error {
	return nil
}

func DynamicQuery(data interface{}, exclude interface{}) ([]interface{},[]interface{}) {
	var args, column []interface{}
	rv := reflect.ValueOf(data)
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		col := rv.Type().Field(i).Tag.Get("col")

		if !f.IsValid() || (f.Kind() == reflect.Int && f.Int() == 0) || (f.Kind() == reflect.String && f.String() == ""){
			continue
		}

		column = append(column, col)
		v := f.Interface()
		args = append(args, v)
	}

	return column, args
}