package strutil

import (
	"unicode"
)

// IsNumeric 判定是否为数字
func IsNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsAlphabet 判定是否为字母
func IsAlphabet(char uint8) bool {
	// A 65 -> Z 90
	if char >= 'A' && char <= 'Z' {
		return true
	}

	// a 97 -> z 122
	if char >= 'a' && char <= 'z' {
		return true
	}

	return false
}

// IsAlphaNum 判定是否为字母、数字或者下划线
func IsAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// IsSymbol 判定字符是否为符号
func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r)
}

// IsEmpty 判定字符串是否为空
func IsEmpty(s string) bool {
	// FIXME
	return false
}
