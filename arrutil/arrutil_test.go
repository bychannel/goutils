package arrutil_test

import (
	"github.com/bychannel/goutils/arrutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ss := []string{"a", "b", "c"}
	arrutil.Reverse(ss)
	assert.Equal(t, []string{"c", "b", "a"}, ss)

	ss2 := []string{"a", "b", "c", "d"}
	arrutil.Reverse(ss2)
	assert.Equal(t, []string{"d", "c", "b", "a"}, ss2)
}

func TestStringsRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}
	ns := arrutil.Remove(ss, "b")

	assert.Contains(t, ns, "a")
	assert.NotContains(t, ns, "b")
	assert.Len(t, ns, 2)
}
