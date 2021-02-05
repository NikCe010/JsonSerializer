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
			part := fmt.Sprintf("\"%s\":\"%s\"", t.Field(i).Name, f.Interface())
			sb = append(sb, part)
		} else if f.Kind() == reflect.Struct {
			sb = append(sb, fmt.Sprintf("\"%s\"", t.Field(i).Name), ":")
			sb = append(sb, serialize(f.Interface())...)
		} else if f.Kind() == reflect.Int {
			part := fmt.Sprintf("\"%s\":%d", t.Field(i).Name, f.Int())
			sb = append(sb, part)
		} else if f.Kind() == reflect.Float32 || f.Kind() == reflect.Float64 {
			part := fmt.Sprintf("\"%s\":%f", t.Field(i).Name, f.Float())
			sb = append(sb, part)
		} else if f.Kind() == reflect.Bool {
			part := fmt.Sprintf("\"%s\":%t", t.Field(i).Name, f.Bool())
			sb = append(sb, part)
		}
		if i != s.NumField()-1 {
			sb = append(sb, ",")
		}
	}
	sb = append(sb, "}")
	return sb
}

//Serialize model to json.
//Should return valid json.
func Serialize(model interface{}) string {
	res := serialize(model)
	return strings.Join(res, "")
}
