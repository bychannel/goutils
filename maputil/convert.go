package maputil

import (
	"errors"
	"github.com/bychannel/goutils/stringutil"
	"reflect"
	"strings"
)

// HttpQueryString convert map[string]interface{} data to httputil query string.
func HttpQueryString(data map[string]interface{}) string {
	ss := make([]string, 0, len(data))
	for k, v := range data {
		ss = append(ss, k+"="+stringutil.QuietString(v))
	}

	return strings.Join(ss, "&")
}

// StructToMap simple convert structs to map by reflect
func StructToMap(st interface{}) (map[string]interface{}, error) {
	mp := make(map[string]interface{})
	if st == nil {
		return mp, nil
	}

	obj := reflect.ValueOf(st)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}

	refType := obj.Type()
	if refType.Kind() != reflect.Struct {
		return mp, errors.New("must be an struct")
	}

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		if field.CanInterface() {
			mp[refType.Field(i).Name] = field.Interface()
		}
	}

	return mp, nil
}
