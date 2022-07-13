package maputil_test

import (
	"fmt"
	"github.com/bychannel/goutils/maputil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpQueryString(t *testing.T) {
	src := map[string]interface{}{"a": "v0", "b": 23}
	str := maputil.HttpQueryString(src)

	fmt.Println(str)
	assert.Contains(t, str, "b=23")
	assert.Contains(t, str, "a=v0")
}
