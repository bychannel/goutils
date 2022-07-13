package httpreq_test

import (
	"encoding/json"
	"fmt"
	"github.com/bychannel/goutils/netutil/httpctype"
	"github.com/bychannel/goutils/netutil/httpreq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpReq_Send(t *testing.T) {
	resp, err := httpreq.New("https://httpbin.org").
		StringBody("hi").
		ContentType(httpctype.JSON).
		Send("/get")

	assert.NoError(t, err)
	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
	assert.False(t, httpreq.IsRedirect(sc))
	assert.False(t, httpreq.IsForbidden(sc))
	assert.False(t, httpreq.IsNotFound(sc))
	assert.False(t, httpreq.IsClientError(sc))
	assert.False(t, httpreq.IsServerError(sc))

	retMp := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&retMp)
	assert.NoError(t, err)
	fmt.Println(retMp)
}

func TestHttpReq_MustSend(t *testing.T) {
	resp := httpreq.New("https://httpbin.org").
		BytesBody([]byte("hi")).
		Method("POST").
		MustSend("/post")

	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
}
