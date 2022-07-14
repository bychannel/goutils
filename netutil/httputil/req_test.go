package httputil_test

import (
	"encoding/json"
	"fmt"
	"github.com/bychannel/goutils/netutil/httputil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpReq_Send(t *testing.T) {
	resp, err := httputil.New("https://httpbin.org").
		StringBody("hi").
		ContentType(httputil.JSON).
		Send("/get")

	assert.NoError(t, err)
	sc := resp.StatusCode
	assert.True(t, httputil.IsOK(sc))
	assert.True(t, httputil.IsSuccessful(sc))
	assert.False(t, httputil.IsRedirect(sc))
	assert.False(t, httputil.IsForbidden(sc))
	assert.False(t, httputil.IsNotFound(sc))
	assert.False(t, httputil.IsClientError(sc))
	assert.False(t, httputil.IsServerError(sc))

	retMp := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&retMp)
	assert.NoError(t, err)
	fmt.Println(retMp)
}

func TestHttpReq_MustSend(t *testing.T) {
	resp := httputil.New("https://httpbin.org").
		BytesBody([]byte("hi")).
		Method("POST").
		MustSend("/post")

	sc := resp.StatusCode
	assert.True(t, httputil.IsOK(sc))
	assert.True(t, httputil.IsSuccessful(sc))
}
