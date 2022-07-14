package mathutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	is := assert.New(t)

	// To int64
	i64Val, err := ToInt64("2")
	is.Nil(err)
	is.Equal(int64(2), i64Val)
}

func TestToFloat(t *testing.T) {
	is := assert.New(t)

	fltVal, err := ToFloat("123.5")
	is.Nil(err)
	is.Equal(123.5, fltVal)
}
