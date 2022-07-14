package arrutil_test

import (
	"fmt"
	"github.com/bychannel/goutils/arrutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt64s(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.ToInt64s([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.SliceToInt64s([]interface{}{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.ToInt64s([]string{"a", "b"})
	is.Error(err)
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.StringsToInts([]string{"a", "b"})
	is.Error(err)
}
