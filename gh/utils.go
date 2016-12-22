package gh

import (
	"fmt"
	"reflect"
	"strconv"
)

func mustString(s interface{}) string {
	if s == nil {
		return ""
	}
	switch t := reflect.ValueOf(s); t.Kind() {
	case reflect.String:
		return s.(string)
	case reflect.Ptr:
		if t.IsNil() {
			return ""
		}

		switch t.Elem().Kind() {
		case reflect.String:
			return *(s.(*string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.Itoa(*(s.(*int)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return fmt.Sprint(*(s.(*uint)))
		case reflect.Float32, reflect.Float64:
			return fmt.Sprint(*(s.(*float64)))
		default:
			return ""
		}
	default:
		return fmt.Sprint(s)
	}
}
