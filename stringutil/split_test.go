package stringutil_test

import (
	"github.com/bychannel/goutils/stringutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", stringutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", stringutil.Substr("abcDef", 2, 2))
	assert.Equal(t, "cDef", stringutil.Substr("abcDef", 2, 0))
	assert.Equal(t, "", stringutil.Substr("abcDEF", 23, 5))
	assert.Equal(t, "cDEF12", stringutil.Substr("abcDEF123", 2, -1))
	assert.Equal(t, "cDEF", stringutil.Substr("abcDEF123", 2, -3))
}
