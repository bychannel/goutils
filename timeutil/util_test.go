package timeutil_test

import (
	"testing"
	"time"

	"github.com/bychannel/goutils/timeutil"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	sec := timeutil.NowUnix()

	assert.NotEmpty(t, timeutil.FormatUnix(sec))
	assert.NotEmpty(t, timeutil.FormatUnixBy(sec, time.RFC3339))

	tt := timeutil.TodayStart()
	assert.Equal(t, "00:00:00", timeutil.DateFormat(tt, "H:I:S"))

	tt = timeutil.TodayEnd()
	assert.Equal(t, "23:59:59", timeutil.DateFormat(tt, "H:I:S"))

	tt = timeutil.NowHourStart()
	assert.Equal(t, "00:00", timeutil.DateFormat(tt, "I:S"))

	tt = timeutil.NowHourEnd()
	assert.Equal(t, "59:59", timeutil.DateFormat(tt, "I:S"))
}

func TestDateFormat(t *testing.T) {
	now := time.Now()

	tests := []struct{ layout, template string }{
		{"20060102 15:04:05", "Ymd H:I:S"},
		{"2006-01-02 15:04:05", "Y-m-d H:I:S"},
		{"2006-01-02 15:04", "Y-m-d H:I"},
		{"01/02 15:04:05", "m/d H:I:S"},
		{"06/01/02 15:04:05", "y/m/d H:I:S"},
		{"06/01/02 15:04:05.000", "y/m/d H:I:Sv"},
	}

	for i, item := range tests {
		date := now.Format(item.layout)
		assert.Equal(t, date, timeutil.DateFormat(now, item.template))
		if i%2 == 0 {
			assert.Equal(t, date, timeutil.Date(now, item.template))
		}
	}

	assert.Equal(t, now.Format("01/02 15:04:05.000000"), timeutil.Date(now, "m/d H:I:Su"))
}

func TestFormatUnix(t *testing.T) {
	now := time.Now()
	want := now.Format("2006-01-02 15:04:05")

	assert.Equal(t, want, timeutil.FormatUnix(now.Unix()))
	assert.Equal(t, want, timeutil.FormatUnixBy(now.Unix(), timeutil.DefaultLayout))
	assert.Equal(t, want, timeutil.FormatUnixByTpl(now.Unix(), "Y-m-d H:I:S"))
}

func TestToLayout(t *testing.T) {
	assert.Equal(t, timeutil.DefaultLayout, timeutil.ToLayout(""))
	assert.Equal(t, time.RFC3339, timeutil.ToLayout("c"))
	assert.Equal(t, time.RFC3339, timeutil.ToLayout("Y-m-dTH:I:SP"))
}
