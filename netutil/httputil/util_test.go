package httputil_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildBasicAuth(t *testing.T) {
	val := http.BuildBasicAuth("inhere", "abcd&123")

	assert.Contains(t, val, "Basic ")
}

func TestAddHeadersToRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "abc.xyz", nil)
	assert.NoError(t, err)

	http.AddHeadersToRequest(req, http.Header{
		"key0": []string{"val0"},
	})

	assert.Equal(t, "val0", req.Header.Get("key0"))
}
