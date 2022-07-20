package mathutil

// MaxInt compare and return max value
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MaxI64 compare and return max value
func MaxI64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// AddI32Exact 安全的计算两个数之和
// 如果溢出了就返回int32最大值
func AddI32Exact(x, y int32) int32 {
	// FIXME
	return 0
}

// AddI64Exact 安全的计算两个数之和
// 如果溢出了就返回int64最大值
func AddI64Exact(x, y int64) int64 {
	// FIXME
	return 0
}

// MultiplyI32Exact 安全计算两个数的乘积
// 如果溢出了就返回int32最大值
func MultiplyI32Exact(x, y int32) int32 {
	// fixme
	return 0
}

// MultiplyI64Exact 安全计算两个数的乘积
// 如果溢出了就返回int64最大值
func MultiplyI64Exact(x, y int64) int64 {
	// fixme
	return 0
}

// Ceil32 向上取整
func Ceil32(a float32) int32 {
	// fixme
	return 0
}

// Ceil64 向上取整
func Ceil64(a float64) int64 {
	// fixme
	return 0
}

// Floor32 向下取整
func Floor32(a float32) int32 {
	// fixme
	return 0
}

// Floor64 向下取整
func Floor64(a float64) int64 {
	// fixme
	return 0
}

// Round32 四舍五入取整
func Round32(a float32) int32 {
	// fixme
	return 0
}

// Round64 四舍五入取整
func Round64(a float64) int64 {
	// fixme
	return 0
}
