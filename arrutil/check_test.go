package arrutil_test

import (
	"github.com/bychannel/goutils/arrutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntsHas(t *testing.T) {
	ints := []int{2, 4, 5}
	assert.True(t, arrutil.IntsHas(ints, 2))
	assert.True(t, arrutil.IntsHas(ints, 5))
	assert.False(t, arrutil.IntsHas(ints, 3))
}

func TestInt64sHas(t *testing.T) {
	ints := []int64{2, 4, 5}
	assert.True(t, arrutil.Int64sHas(ints, 2))
	assert.True(t, arrutil.Int64sHas(ints, 5))
	assert.False(t, arrutil.Int64sHas(ints, 3))
}

func TestStringsHas(t *testing.T) {
	ss := []string{"a", "b"}
	assert.True(t, arrutil.StringsHas(ss, "a"))
	assert.True(t, arrutil.StringsHas(ss, "b"))

	assert.False(t, arrutil.StringsHas(ss, "c"))
}
