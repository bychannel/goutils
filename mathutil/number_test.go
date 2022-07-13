package mathutil_test

import (
	"testing"
	"time"

	"github.com/bychannel/goutils/mathutil"
	"github.com/stretchr/testify/assert"
)

func TestIsNumeric(t *testing.T) {
	assert.True(t, mathutil.IsNumeric('3'))
	assert.False(t, mathutil.IsNumeric('a'))
}

func TestPercent(t *testing.T) {
	assert.Equal(t, float64(34), mathutil.Percent(34, 100))
	assert.Equal(t, float64(0), mathutil.Percent(34, 0))
	assert.Equal(t, float64(-100), mathutil.Percent(34, -34))
}

func TestElapsedTime(t *testing.T) {
	nt := time.Now().Add(-time.Second * 3)
	num := mathutil.ElapsedTime(nt)

	assert.Equal(t, 3000, int(mathutil.MustFloat(num)))
}
