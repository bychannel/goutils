package httputil_test

import (
	"github.com/bychannel/goutils/netutil/httputil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHeaderToStringMap(t *testing.T) {
	assert.Nil(t, httputil.HeaderToStringMap(nil))
	assert.Nil(t, httputil.HeaderToStringMap(http.Header{}))

	want := map[string]string{"key": "value; more"}
	assert.Equal(t, want, httputil.HeaderToStringMap(http.Header{
		"key": {"value", "more"},
	}))
}
