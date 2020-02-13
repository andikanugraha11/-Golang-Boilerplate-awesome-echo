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
		if !f.IsValid() || f.String() == "" || f.Int() == 0{
			continue
		}

		v := f.Interface()
		args = append(args, v)
		column = append(column, col)
	}

	return args, column
}