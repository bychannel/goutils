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

func TestEquals(t *testing.T) {
	assert.True(t, stringutil.Equal("a", "a"))
	assert.False(t, stringutil.Equal("a", "b"))
}

func TestLen(t *testing.T) {
	str := "Hello, 世界"

	assert.Equal(t, 7, stringutil.Len("Hello, "))
	assert.Equal(t, 13, stringutil.Len(str))
	assert.Equal(t, 9, stringutil.RuneLen(str))
	assert.Equal(t, 9, stringutil.Utf8len(str))
	assert.Equal(t, 9, stringutil.Utf8Len(str))
	assert.True(t, stringutil.IsValidUtf8(str))
}

func TestStrPos(t *testing.T) {
	// StrPos
	assert.Equal(t, -1, stringutil.StrPos("xyz", "a"))
	assert.Equal(t, 0, stringutil.StrPos("xyz", "x"))
	assert.Equal(t, 2, stringutil.StrPos("xyz", "z"))

	// RunePos
	assert.Equal(t, -1, stringutil.RunePos("xyz", 'a'))
	assert.Equal(t, 0, stringutil.RunePos("xyz", 'x'))
	assert.Equal(t, 2, stringutil.RunePos("xyz", 'z'))
	assert.Equal(t, 5, stringutil.RunePos("hi时间", '间'))

	// BytePos
	assert.Equal(t, -1, stringutil.BytePos("xyz", 'a'))
	assert.Equal(t, 0, stringutil.BytePos("xyz", 'x'))
	assert.Equal(t, 2, stringutil.BytePos("xyz", 'z'))
	// assert.Equal(t, 2, stringutil.BytePos("hi时间", '间')) // will build error
}

func TestIsStartOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "a", true},
		{"abc", "d", false},
	}

	for _, item := range tests {
		assert.Equal(t, item.want, stringutil.HasPrefix(item.give, item.sub))
		assert.Equal(t, item.want, stringutil.IsStartOf(item.give, item.sub))
	}

	assert.True(t, stringutil.IsStartsOf("abc", []string{"a", "b"}))
	assert.False(t, stringutil.IsStartsOf("abc", []string{"d", "e"}))
}

func TestIsEndOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "c", true},
		{"abc", "d", false},
		{"some.json", ".json", true},
	}

	for _, item := range tests {
		assert.Equal(t, item.want, stringutil.HasSuffix(item.give, item.sub))
		assert.Equal(t, item.want, stringutil.IsEndOf(item.give, item.sub))
	}
}

func TestIsSpace(t *testing.T) {
	assert.True(t, stringutil.IsSpace(' '))
	assert.True(t, stringutil.IsSpace('\n'))
	assert.True(t, stringutil.IsSpaceRune('\n'))
	assert.True(t, stringutil.IsSpaceRune('\t'))

	assert.False(t, stringutil.IsBlank(" a "))
	assert.True(t, stringutil.IsNotBlank(" a "))
	assert.False(t, stringutil.IsEmpty(" "))
	assert.True(t, stringutil.IsBlank(" "))
	assert.True(t, stringutil.IsBlank("   "))
	assert.False(t, stringutil.IsNotBlank("   "))

	assert.False(t, stringutil.IsBlankBytes([]byte(" a ")))
	assert.True(t, stringutil.IsBlankBytes([]byte(" ")))
	assert.True(t, stringutil.IsBlankBytes([]byte("   ")))
}

func TestIsSymbol(t *testing.T) {
	assert.False(t, stringutil.IsSymbol('a'))
	assert.True(t, stringutil.IsSymbol('●'))
}

func TestHasOneSub(t *testing.T) {
	assert.False(t, stringutil.HasOneSub("h3ab2c", []string{"d"}))
	assert.True(t, stringutil.HasOneSub("h3ab2c", []string{"ab"}))
}

func TestHasAllSubs(t *testing.T) {
	assert.False(t, stringutil.HasAllSubs("h3ab2c", []string{"a", "d"}))
	assert.True(t, stringutil.HasAllSubs("h3ab2c", []string{"a", "b"}))
}

func TestVersionCompare(t *testing.T) {
	versions := []struct{ a, b string }{
		{"1.0.221.9289", "1.05.00.0156"},
		// Go versions
		{"1", "1.0.1"},
		{"1.0.1", "1.0.2"},
		{"1.0.2", "1.0.3"},
		{"1.0.3", "1.1"},
		{"1.1", "1.1.1"},
		{"1.1.1", "1.1.2"},
		{"1.1.2", "1.2"},
	}
	for _, version := range versions {
		assert.True(t, stringutil.VersionCompare(version.a, version.b, "<"), version.a+"<"+version.b)
		assert.True(t, stringutil.VersionCompare(version.a, version.b, "<="), version.a+"<="+version.b)
		assert.True(t, stringutil.VersionCompare(version.b, version.a, ">"), version.a+">"+version.b)
		assert.True(t, stringutil.VersionCompare(version.b, version.a, ">="), version.a+">="+version.b)
	}

	assert.True(t, stringutil.VersionCompare("1.0", "1.0", ""))
	assert.True(t, stringutil.VersionCompare("1.0", "1.0", "="))
}
