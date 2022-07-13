package stringutil_test

import (
	"github.com/bychannel/goutils/stringutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlphabet(t *testing.T) {
	assert.True(t, stringutil.IsNumeric('9'))
	assert.False(t, stringutil.IsNumeric('A'))

	assert.False(t, stringutil.IsAlphabet('9'))
	assert.False(t, stringutil.IsAlphabet('+'))

	assert.True(t, stringutil.IsAlphabet('A'))
	assert.True(t, stringutil.IsAlphabet('a'))
	assert.True(t, stringutil.IsAlphabet('Z'))
	assert.True(t, stringutil.IsAlphabet('z'))
}

func TestIsAlphaNum(t *testing.T) {
	assert.False(t, stringutil.IsAlphaNum('+'))

	assert.True(t, stringutil.IsAlphaNum('9'))
	assert.True(t, stringutil.IsAlphaNum('A'))
	assert.True(t, stringutil.IsAlphaNum('a'))
	assert.True(t, stringutil.IsAlphaNum('Z'))
	assert.True(t, stringutil.IsAlphaNum('z'))
}

func TestIsSymbol(t *testing.T) {
	assert.False(t, stringutil.IsSymbol('a'))
	assert.True(t, stringutil.IsSymbol('●'))
}
