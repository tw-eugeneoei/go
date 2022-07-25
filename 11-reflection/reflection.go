package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	// * v1
	// value := getValue(x)

	// if value.Kind() == reflect.Slice {
	// 	for i := 0; i < value.Len(); i++ {
	// 		Walk(value.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// for i := 0; i < value.NumField(); i++ {
	// 	field := value.Field(i)

	// 	switch field.Kind() {
	// 	case reflect.String:
	// 		fn(field.String())
	// 	case reflect.Struct:
	// 		Walk(field.Interface(), fn)
	// 	case reflect.Slice:
	// 		for i := 0; i < value.Len(); i++ {
	// 			Walk(value.Index(i).Interface(), fn)
	// 		}
	// 	}
	// }

	// * v2
	value := getValue(x)
	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}
	switch value.Kind() {

	case reflect.String:
		fn(value.String())

	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}

	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	}

	// // * v3
	// value := getValue(x)
	// numberOfValues := 0
	// var getField func(int) reflect.Value

	// switch value.Kind() {

	// case reflect.String:
	// 	fn(value.String())

	// case reflect.Struct:
	// 	numberOfValues = value.NumField()
	// 	getField = value.Field

	// case reflect.Slice, reflect.Array:
	// 	numberOfValues = value.Len()
	// 	getField = value.Index

	// case reflect.Map:
	// 	// fmt.Println(value.MapKeys())
	// 	// fmt.Println(value.Len())
	// 	// for _, key := range value.MapKeys() {
	// 	// 	Walk(value.MapIndex(key).Interface(), fn)
	// 	// }
	// 	numberOfValues = value.Len()
	// 	getField = value.MapIndex

	// }

	// for i := 0; i < numberOfValues; i++ {
	// 	Walk(getField(i).Interface(), fn)
	// }
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	// * You can't use NumField on a pointer Value, we need to extract the underlying value
	// * before we can do that by using Elem().
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}
