package strutil_test

import (
	"github.com/bychannel/goutils/strutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", strutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", strutil.Substr("abcDef", 2, 2))
	assert.Equal(t, "cDef", strutil.Substr("abcDef", 2, 0))
	assert.Equal(t, "", strutil.Substr("abcDEF", 23, 5))
	assert.Equal(t, "cDEF12", strutil.Substr("abcDEF123", 2, -1))
	assert.Equal(t, "cDEF", strutil.Substr("abcDEF123", 2, -3))
}
