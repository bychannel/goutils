package stringutil_test

import (
	"github.com/bychannel/goutils/stringutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", stringutil.Md5("123456"))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", stringutil.MD5("123456"))
	assert.Equal(t, "a906449d5769fa7361d7ecc6aa3f6d28", stringutil.GenMd5("123abc"))
	assert.Equal(t, "289dff07669d7a23de0ef88d2f7129e7", stringutil.GenMd5(234))
}

func TestEscape(t *testing.T) {
	tests := struct{ give, want string }{
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Equal(t, tests.want, stringutil.EscapeHTML(tests.give))

	ret := stringutil.EscapeJS("<script>var a = 23;</script>")
	assert.NotContains(t, ret, "<script>")
	assert.NotContains(t, ret, "</script>")
}

func TestAddSlashes(t *testing.T) {
	assert.Equal(t, "", stringutil.AddSlashes(""))
	assert.Equal(t, "", stringutil.StripSlashes(""))

	assert.Equal(t, `{\"key\": 123}`, stringutil.AddSlashes(`{"key": 123}`))
	assert.Equal(t, `{"key": 123}`, stringutil.StripSlashes(`{\"key\": 123}`))
}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", stringutil.URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", stringutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", stringutil.URLEncode("a.com"))
	is.Equal("a.com", stringutil.URLDecode("a.com"))
}
