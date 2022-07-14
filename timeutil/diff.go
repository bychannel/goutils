package timeutil

import "time"

// DiffSeconds 计算两个时间点的秒数差
func DiffSeconds(t, u time.Time) int32 {
	return int32(t.Sub(u) / time.Second)
}

// DiffMinutes 计算两个时间点的分钟差
func DiffMinutes(t, u time.Time) int32 {
	return int32(t.Sub(u) / time.Minute)
}

// DiffHours 计算两个时间点的小时差
func DiffHours(t, u time.Time) int32 {
	return int32(t.Sub(u) / time.Hour)
}
