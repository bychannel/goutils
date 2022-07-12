package stringutil_test

import (
	"github.com/bychannel/goutils/stringutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	is := assert.New(t)

	// Trim
	tests := map[string]string{
		"abc ":  "",
		" abc":  "",
		" abc ": "",
		"abc,,": ",",
		"abc,.": ",.",
	}
	for sample, cutSet := range tests {
		is.Equal("abc", stringutil.Trim(sample, cutSet))
	}

	is.Equal("abc", stringutil.Trim("abc,.", ".,"))
	is.Equal("abc", stringutil.Trim(", abc ,", ",", " "))

	// TrimLeft
	is.Equal("abc ", stringutil.Ltrim(" abc "))
	is.Equal("abc ", stringutil.LTrim(" abc "))
	is.Equal("abc ,", stringutil.TrimLeft(", abc ,", " ,"))
	is.Equal("abc ,", stringutil.TrimLeft(", abc ,", ", "))
	is.Equal("abc ,", stringutil.TrimLeft(", abc ,", ",", " "))
	is.Equal(" abc ,", stringutil.TrimLeft(", abc ,", ","))

	// TrimRight
	is.Equal(" abc", stringutil.Rtrim(" abc "))
	is.Equal(" abc", stringutil.RTrim(" abc "))
	is.Equal(", abc", stringutil.TrimRight(", abc ,", ", "))
	is.Equal(", abc ", stringutil.TrimRight(", abc ,", ","))
}

func TestFilterEmail(t *testing.T) {
	is := assert.New(t)
	is.Equal("THE@inhere.com", stringutil.FilterEmail("   THE@INHere.com  "))
	is.Equal("inhere.xyz", stringutil.FilterEmail("   inhere.xyz  "))
}
