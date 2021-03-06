package maputil

import (
	"reflect"
	"strconv"
	"strings"
)

// GetOrDefault 获取给定key的值，如果没有则返回默认值
func GetOrDefault(m map[string]interface{}, key string, def interface{}) interface{} {
	val, ok := m[key]
	if ok {
		return val
	}
	return def
}

// GetOrMaxKey 从给定map中获取指定的key，如果key不存在，那就获取最大的key所对应的值
func GetOrMaxKey(m map[int32]interface{}, key int32) interface{} {
	// FIXME
	return nil
}

// AddByI32Value 将集合m2加到集合m1上
func AddByI32Value(m1, m2 map[interface{}]int32) map[interface{}]int32 {
	// FIXME
	return nil
}

// AddByI64Value 将集合m2加到集合m1上
func AddByI64Value(m1, m2 map[interface{}]int64) map[interface{}]int64 {
	// FIXME
	return nil
}

// MergeStringMap simple merge two string map. merge src to dst map
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}

		dst[k] = v
	}
	return dst
}

// GetByPath get value from a map[string]interface{}. eg "top" "top.sub"
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool) {
	if val, ok := mp[key]; ok {
		return val, true
	}

	// has sub key? eg. "top.sub"
	if !strings.ContainsRune(key, '.') {
		return nil, false
	}

	keys := strings.Split(key, ".")
	topK := keys[0]

	// find top item data based on top key
	var item interface{}
	if item, ok = mp[topK]; !ok {
		return
	}

	for _, k := range keys[1:] {
		switch tData := item.(type) {
		case map[string]string: // is simple map
			if item, ok = tData[k]; !ok {
				return
			}
		case map[string]interface{}: // is map(decode from toml/json)
			if item, ok = tData[k]; !ok {
				return
			}
		case map[interface{}]interface{}: // is map(decode from yaml)
			if item, ok = tData[k]; !ok {
				return
			}
		case []interface{}: // is a slice
			if item, ok = getBySlice(k, tData); !ok {
				return
			}
		case []string, []int, []float32, []float64, []bool, []rune:
			slice := reflect.ValueOf(tData)
			sData := make([]interface{}, slice.Len())
			for i := 0; i < slice.Len(); i++ {
				sData[i] = slice.Index(i).Interface()
			}
			if item, ok = getBySlice(k, sData); !ok {
				return
			}
		default: // error
			return nil, false
		}
	}

	return item, true
}

// Keys get all keys of the given map.
func Keys(mp interface{}) (keys []string) {
	rftVal := reflect.Indirect(reflect.ValueOf(mp))
	if rftVal.Kind() != reflect.Map {
		return
	}

	keys = make([]string, 0, rftVal.Len())
	for _, key := range rftVal.MapKeys() {
		keys = append(keys, key.String())
	}
	return
}

// Values get all values from the given map.
func Values(mp interface{}) (values []interface{}) {
	rftVal := reflect.Indirect(reflect.ValueOf(mp))
	if rftVal.Kind() != reflect.Map {
		return
	}

	values = make([]interface{}, 0, rftVal.Len())
	for _, key := range rftVal.MapKeys() {
		values = append(values, rftVal.MapIndex(key).Interface())
	}
	return
}

func getBySlice(k string, slice []interface{}) (val interface{}, ok bool) {
	i, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return nil, false
	}
	if size := int64(len(slice)); i >= size {
		return nil, false
	}
	return slice[i], true
}
