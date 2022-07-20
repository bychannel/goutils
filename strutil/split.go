package strutil

// Substr for a string.
// if length <= 0, return pos to end.
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	strLn := len(runes)

	// pos is too large
	if pos >= strLn {
		return ""
	}

	stopIdx := pos + length
	if length == 0 || stopIdx > strLn {
		stopIdx = strLn
	} else if length < 0 {
		stopIdx = strLn + length
	}

	return string(runes[pos:stopIdx])
}
