package timeutil

import (
	"time"
)

var DefaultLayout = "2006-01-02 15:04:05"

// Format 默认形式格式化时间
func Format(t time.Time) string {
	return t.Format(DefaultLayout)
}

// FormatBy 给定形式格式化时间
func FormatBy(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatUnix 默认形式格式化时间戳
func FormatUnix(sec int64) string {
	return time.Unix(sec, 0).Format(DefaultLayout)
}

// FormatUnixBy 给定形式格式化时间戳
func FormatUnixBy(sec int64, layout string) string {
	return time.Unix(sec, 0).Format(layout)
}

// CustomHMS 自定义时、分、秒
func CustomHMS(hour, min, sec int) time.Time {
	now := time.Now()
	y, m, d := now.Date()
	return time.Date(y, m, d, hour, min, sec, int(time.Second-time.Nanosecond), now.Location())
}

// IsSameDay 判断给定的两个时间是否同一天
func IsSameDay(d1, d2 time.Time) bool {
	// fixme
	return false
}

// IsSameDayWithOffset 判断给定的两个时间是否同一天
// offset为秒钟偏移值
// 例如：需要在次日的5点刷新，offset传值为 5*3600 = 18000
func IsSameDayWithOffset(d1, d2 time.Time, offset int32) bool {
	// fixme
	return false
}
