package maputil

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gookit/goutil/strutil"
)

// KeyToLower convert keys to lower case.
func KeyToLower(src map[string]string) map[string]string {
	newMp := make(map[string]string, len(src))
	for k, v := range src {
		k = strings.ToLower(k)
		newMp[k] = v
	}

	return newMp
}

// ToStringMap convert map[string]interface{} to map[string]string
func ToStringMap(src map[string]interface{}) map[string]string {
	newMp := make(map[string]string, len(src))
	for k, v := range src {
		newMp[k] = strutil.MustString(v)
	}

	return newMp
}

// HttpQueryString convert map[string]interface{} data to http query string.
func HttpQueryString(data map[string]interface{}) string {
	ss := make([]string, 0, len(data))
	for k, v := range data {
		ss = append(ss, k+"="+strutil.QuietString(v))
	}

	return strings.Join(ss, "&")
}

// ToString simple and quickly convert map[string]interface{} to string.
func ToString(mp map[string]interface{}) string {
	if mp == nil {
		return ""
	}
	if len(mp) == 0 {
		return "{}"
	}

	buf := make([]byte, 0, len(mp)*16)
	buf = append(buf, '{')

	for k, val := range mp {
		buf = append(buf, k...)
		buf = append(buf, ':')

		str := strutil.QuietString(val)
		buf = append(buf, str...)
		buf = append(buf, ',', ' ')
	}

	// remove last ', '
	buf = append(buf[:len(buf)-2], '}')
	return strutil.Byte2str(buf)
}

// ToString2 simple and quickly convert a map to string.
func ToString2(mp interface{}) string {
	return NewFormatter(mp).Format()
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
