package timeutil_test

import (
	"github.com/bychannel/goutils/timeutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeX_basic(t *testing.T) {
	tx := timeutil.Now()
	assert.NotEmpty(t, tx.String())
	assert.NotEmpty(t, tx.Datetime())
}

func TestTimeX_Format(t *testing.T) {
	tx := timeutil.Now()
	assert.Equal(t, tx.Datetime(), tx.DateFormat("Y-m-d H:I:S"))
}

func TestTimeX_SubUnix(t *testing.T) {
	tx := timeutil.Now()

	after1m := tx.AddMinutes(1)

	assert.Equal(t, timeutil.OneMinSec, after1m.SubUnix(tx.Time))
}

func TestTimeX_DateFormat(t *testing.T) {
	tx := timeutil.Now()
	assert.Equal(t, tx.Format(timeutil.DefaultLayout), tx.DateFormat("Y-m-d H:I:S"))
	assert.Equal(t, tx.Format(""), tx.DateFormat("Y-m-d H:I:S"))
	assert.Equal(t, tx.Format("2006/01/02 15:04"), tx.TplFormat("Y/m/d H:I"))

	assert.Equal(t, "23:59:59", tx.DayEnd().DateFormat("H:I:S"))
	assert.Equal(t, "00:00:00", tx.DayStart().DateFormat("H:I:S"))
}

func TestTimeX_AddDay(t *testing.T) {
	tx := timeutil.Now()

	yd := tx.Yesterday()
	yd1 := tx.AddDay(-1)
	assert.Equal(t, yd.Unix(), yd1.Unix())
	assert.Equal(t, yd.Unix(), tx.DayAgo(1).Unix())

	assert.True(t, tx.IsAfter(yd.Time))
	assert.True(t, tx.IsAfterUnix(yd.Time.Unix()))
	assert.True(t, yd.IsBefore(tx.Time))
	assert.True(t, yd.IsBeforeUnix(tx.T().Unix()))

	assert.Equal(t, tx.Unix()-yd.Unix(), int64(timeutil.OneDaySec))

	md := tx.Tomorrow()
	yd2 := tx.DayAfter(1)
	assert.Equal(t, md.Unix(), yd2.Unix())
}

func TestTimeX_AddSeconds(t *testing.T) {
	tx := timeutil.Now()

	h1 := tx.AddHour(1)
	s1 := tx.AddSeconds(timeutil.OneHourSec)
	assert.Equal(t, h1.Unix(), s1.Unix())

	assert.Equal(t, timeutil.OneHour, h1.Diff(tx.Time))
	assert.Equal(t, timeutil.OneHourSec, h1.DiffSec(tx.Time))
}

func TestTimeX_HourStart(t *testing.T) {
	tx := timeutil.Now()
	hs := tx.HourStart()
	he := tx.HourEnd()

	assert.Equal(t, "00:00", hs.DateFormat("I:S"))
	assert.Equal(t, "59:59", he.DateFormat("I:S"))
}

func TestTimeX_CustomHMS(t *testing.T) {
	tx := timeutil.Now()
	assert.Equal(t, "12:23:34", tx.CustomHMS(12, 23, 34).TplFormat("H:I:S"))
}
