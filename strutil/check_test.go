package strutil_test

import (
	"github.com/bychannel/goutils/strutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlphabet(t *testing.T) {
	assert.True(t, strutil.IsNumeric('9'))
	assert.False(t, strutil.IsNumeric('A'))

	assert.False(t, strutil.IsAlphabet('9'))
	assert.False(t, strutil.IsAlphabet('+'))

	assert.True(t, strutil.IsAlphabet('A'))
	assert.True(t, strutil.IsAlphabet('a'))
	assert.True(t, strutil.IsAlphabet('Z'))
	assert.True(t, strutil.IsAlphabet('z'))
}

func TestIsAlphaNum(t *testing.T) {
	assert.False(t, strutil.IsAlphaNum('+'))

	assert.True(t, strutil.IsAlphaNum('9'))
	assert.True(t, strutil.IsAlphaNum('A'))
	assert.True(t, strutil.IsAlphaNum('a'))
	assert.True(t, strutil.IsAlphaNum('Z'))
	assert.True(t, strutil.IsAlphaNum('z'))
}

func TestIsSymbol(t *testing.T) {
	assert.False(t, strutil.IsSymbol('a'))
	assert.True(t, strutil.IsSymbol('●'))
}
