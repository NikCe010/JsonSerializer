package serializer

import (
	"fmt"
	"reflect"
	"strings"
)

func serialize(model interface{}) []string {
	s := reflect.ValueOf(model)
	t := reflect.TypeOf(model)
	var sb []string
	sb = append(sb, "{")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if f.Kind() == reflect.String || f.Type().Name() == "Time" {
			part := fmt.Sprintf("\"%s\":\"%s\"", tryGetTag(t.Field(i)), f.Interface())
			sb = append(sb, part)
		} else if f.Kind() == reflect.Struct {
			sb = append(sb, fmt.Sprintf("\"%s\"", tryGetTag(t.Field(i))), ":")
			sb = append(sb, serialize(f.Interface())...)
		} else if f.Kind() == reflect.Int {
			part := fmt.Sprintf("\"%s\":%d", tryGetTag(t.Field(i)), f.Int())
			sb = append(sb, part)
		} else if f.Kind() == reflect.Float32 || f.Kind() == reflect.Float64 {
			part := fmt.Sprintf("\"%s\":%f", tryGetTag(t.Field(i)), f.Float())
			sb = append(sb, part)
		} else if f.Kind() == reflect.Bool {
			part := fmt.Sprintf("\"%s\":%t", tryGetTag(t.Field(i)), f.Bool())
			sb = append(sb, part)
		}
		if i != s.NumField()-1 {
			sb = append(sb, ",")
		}
	}
	sb = append(sb, "}")
	return sb
}

func tryGetTag(t reflect.StructField) string {
	tag := t.Tag.Get("json")
	if tag == "" {
		return t.Name
	}
	return tag
}

//Serialize model to json.
//Should return valid json.
func Serialize(model interface{}) string {
	res := serialize(model)
	return strings.Join(res, "")
}
