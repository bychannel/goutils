package randomutil

import (
	"math/rand"
	"time"
)

// RandomInt 在[min, max)中获取一个随机值
//
// Usage:
// 	RandomInt(10, 99)
// 	RandomInt(100, 999)
// 	RandomInt(1000, 9999)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomIntWithSeed 在[min, max)中获取一个随机值，自定义随机种子
//
// Usage:
// 	seed := time.Now().UnixNano()
// 	RandomIntWithSeed(1000, 9999, seed)
func RandomIntWithSeed(min, max int, seed int64) int {
	rand.Seed(seed)
	return min + rand.Intn(max-min)
}
