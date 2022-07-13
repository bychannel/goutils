package stringutil_test

import (
	"github.com/bychannel/goutils/stringutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", stringutil.URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", stringutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", stringutil.URLEncode("a.com"))
	is.Equal("a.com", stringutil.URLDecode("a.com"))
}
