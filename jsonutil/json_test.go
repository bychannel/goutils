package jsonutil_test

import (
	"github.com/bychannel/goutils/jsonutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteReadFile(t *testing.T) {
	user := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{"inhere", 200}

	err := jsonutil.WriteFile("testdata/test.json", &user)
	assert.NoError(t, err)

	err = jsonutil.ReadFile("testdata/test.json", &user)
	assert.NoError(t, err)

	assert.Equal(t, "inhere", user.Name)
	assert.Equal(t, 200, user.Age)
}
