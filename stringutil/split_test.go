package stringutil_test

import (
	"fmt"
	"github.com/bychannel/goutils/stringutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	str := "hi,inhere"

	b, a, ok := stringutil.Cut(str, ",")
	assert.True(t, ok)
	assert.Equal(t, "hi", b)
	assert.Equal(t, "inhere", a)

	b, a = stringutil.MustCut(str, ",")
	assert.Equal(t, "hi", b)
	assert.Equal(t, "inhere", a)

	b, a, ok = stringutil.Cut(str, "-")
	assert.False(t, ok)
	assert.Equal(t, str, b)
	assert.Equal(t, "", a)
}

func TestSplit(t *testing.T) {
	ss := stringutil.Split("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = stringutil.SplitValid("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = stringutil.SplitN("a, , b,c", ",", 3)
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = stringutil.SplitN("a, , b,c", ",", 2)
	assert.Equal(t, `[]string{"a", "b,c"}`, fmt.Sprintf("%#v", ss))

	ss = stringutil.SplitN("origin https://github.com/bychannel/goutils (push)", " ", 3)
	assert.Len(t, ss, 3)

	ss = stringutil.Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSplitTrimmed(t *testing.T) {
	ss := stringutil.SplitTrimmed("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = stringutil.SplitNTrimmed("a, , b,c", ",", 2)
	assert.Equal(t, `[]string{"a", ", b,c"}`, fmt.Sprintf("%#v", ss))
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", stringutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", stringutil.Substr("abcDef", 2, 2))
	assert.Equal(t, "cDef", stringutil.Substr("abcDef", 2, 0))
	assert.Equal(t, "", stringutil.Substr("abcDEF", 23, 5))
	assert.Equal(t, "cDEF12", stringutil.Substr("abcDEF123", 2, -1))
	assert.Equal(t, "cDEF", stringutil.Substr("abcDEF123", 2, -3))
}
