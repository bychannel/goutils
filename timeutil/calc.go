package timeutil

import (
	"fmt"
	"time"
)

// ElapsedTime 计算运行时间消耗 单位 ms(毫秒)
func ElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3d", time.Since(startTime).Milliseconds())
}

// AddDay 给定时间点增加天数
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddHour 给定时间点增加小时数
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Duration(hour) * time.Hour)
}

// AddMinutes 给定时间点增加分钟数
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddSeconds 给定时间点增加秒数
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// HourStartBy 计算给定时间点的小时开始时间
func HourStartBy(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// HourEndBy 计算给定时间点的小时结束时间
func HourEndBy(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// DayStartBy 计算给定时间点的当天开始时间
func DayStartBy(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// DayEndBy 计算给定时间点的当天结束时间
func DayEndBy(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}
