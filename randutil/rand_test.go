package randutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	min, max := 1000, 9999

	for i := 0; i < 5; i++ {
		val := RandInt(min, max)
		fmt.Println(val)
		assert.True(t, val >= min)
		assert.True(t, val <= max)

		seed := time.Now().UnixNano()
		val = RandIntWithSeed(min, max, seed)
		assert.True(t, val >= min)
	}
}
