package deserializer

import (
	"reflect"
)

func SetValues(any interface{}, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	setValues(any, inputs...)
}

func setValues(any interface{}, inputs ...reflect.Value) {
	for i := 0; i < len(inputs); i++ {
		field := reflect.ValueOf(any).Elem().Field(i)
		if field.Kind() != reflect.Struct {
			field.Set(inputs[i])
		} else {
			i += setValuesForNestedStruct(field, inputs[i:])
		}
	}
}

func setValuesForNestedStruct(s reflect.Value, inputs []reflect.Value) int {
	i := 0
	for ; i < s.NumField(); i++ {
		field := s.Field(i)
		if field.Kind() != reflect.Struct {
			field.Set(inputs[i])
		} else {
			i += setValuesForNestedStruct(field, inputs[i:])
		}
	}
	return i
}
