package timeutil

import (
	"fmt"
	"time"
)

// Format use default layout
func Format(t time.Time) string {
	return t.Format(DefaultLayout)
}

// FormatBy format time by given date template.
func FormatBy(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatUnix time seconds use default layout
func FormatUnix(sec int64) string {
	return time.Unix(sec, 0).Format(DefaultLayout)
}

// FormatUnixBy format time seconds use given layout
func FormatUnixBy(sec int64, layout string) string {
	return time.Unix(sec, 0).Format(layout)
}

// AddDay 给定时间点增加天数
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddHour 给定时间点增加小时数
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Duration(hour) * OneHour)
}

// AddMinutes 给定时间点增加分钟数
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * OneMin)
}

// AddSeconds 给定时间点增加秒数
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// HourStart 计算给定时间点的小时开始时间
func HourStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// HourEnd 计算给定时间点的小时结束时间
func HourEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// DayStart 计算给定时间点的当天开始时间
func DayStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// DayEnd 计算给定时间点的当天结束时间
func DayEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// ElapsedTime 计算运行时间消耗 单位 ms(毫秒)
func ElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3d", time.Since(startTime).Milliseconds())
}
