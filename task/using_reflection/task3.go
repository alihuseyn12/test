package using_reflection

import (
	"reflect"
)

func SetFieldValue(value interface{}, fieldName string, newValue string) {

	v := reflect.ValueOf(&value).Elem()
	tmp := reflect.New(v.Elem().Type()).Elem()
	tmp.Set(v.Elem())
	tmp.FieldByName("Name").SetString("alik")
	v.Set(tmp)
}
