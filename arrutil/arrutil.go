package arrutil

// Reverse string slice
func Reverse(ss []string) {
	ln := len(ss)
	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		ss[i], ss[li] = ss[li], ss[i]
	}
}

// Remove a value form a string slice
func Remove(ss []string, s string) []string {
	ns := make([]string, 0, len(ss))
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}
	return ns
}
