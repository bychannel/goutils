package strutil_test

import (
	"github.com/bychannel/goutils/strutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", strutil.URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", strutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", strutil.URLEncode("a.com"))
	is.Equal("a.com", strutil.URLDecode("a.com"))
}
