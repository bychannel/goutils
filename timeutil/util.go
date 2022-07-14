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
