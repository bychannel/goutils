package arrutil

import (
	"errors"
	"github.com/bychannel/goutils/mathutil"
	"reflect"
	"strconv"
)

// ErrInvalidType error
var ErrInvalidType = errors.New("the input param type is invalid")

// StringsToInts 字符串切片转整形切片
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

// StringsToSlice 字符串切片转空接口切片
func StringsToSlice(ss []string) []interface{} {
	args := make([]interface{}, len(ss))
	for i, s := range ss {
		args[i] = s
	}
	return args
}

// ToInt64s 切片或数组转int64切片
func ToInt64s(arr interface{}) (ret []int64, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		err = ErrInvalidType
		return
	}

	for i := 0; i < rv.Len(); i++ {
		i64, err := mathutil.Int64(rv.Index(i).Interface())
		if err != nil {
			return []int64{}, err
		}

		ret = append(ret, i64)
	}
	return
}

// SliceToInt64s convert []interface{} to []int64
func SliceToInt64s(arr []interface{}) []int64 {
	i64s := make([]int64, len(arr))
	for i, v := range arr {
		i64s[i] = mathutil.MustInt64(v)
	}
	return i64s
}
