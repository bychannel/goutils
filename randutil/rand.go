package randutil

import (
	"math/rand"
	"time"
)

// RandInt 在[min, max)中获取一个随机值
//
// Usage:
// 	RandInt(10, 99)
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandIntWithSeed 在[min, max)中获取一个随机值，自定义随机种子
//
// Usage:
// 	seed := time.Now().UnixNano()
// 	RandIntWithSeed(1000, 9999, seed)
func RandIntWithSeed(min, max int, seed int64) int {
	rand.Seed(seed)
	return min + rand.Intn(max-min)
}

// RandBool 随机一个bool值
func RandBool() bool {
	// FIXME
	return false
}

// IsSuccess 判断一次随机事件是否成功(比较概率大小)
func IsSuccess(rate float32) bool {
	// FIXME
	return false
}

// IsSuccessByPercentage 判断一次随机事件是否成功(计算百分比大小)
func IsSuccessByPercentage(rate float32) bool {
	// FIXME
	return false
}

// IsSuccessByPermillage 判断一次随机事件是否成功(计算千分比大小)
func IsSuccessByPermillage(rate float32) bool {
	// FIXME
	return false
}

// RandSlice 给定切片中随机出一个元素(无权重)
func RandSlice(arr []interface{}) interface{} {
	// FIXME
	return nil
}

// RandSliceN 给定切片中随机出指定数量个元素(无权重)
func RandSliceN(arr []interface{}, n int32) []interface{} {
	// FIXME
	return nil
}

// RandSliceByWeight 给定集合中随机出一个元素(按权重)
// k为元素，v为权重值
func RandSliceByWeight(arr map[interface{}]int32) interface{} {
	// FIXME
	return nil
}

// RandSliceNByWeight 给定集合中随机出指定数量个元素(按权重)
// k为元素，v为权重值
func RandSliceNByWeight(arr map[interface{}]int32, n int32) []interface{} {
	// FIXME
	return nil
}
