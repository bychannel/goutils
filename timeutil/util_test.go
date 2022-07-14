package timeutil_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/bychannel/goutils/timeutil"
	"github.com/stretchr/testify/assert"
)

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

	for _, item := range tests {
		date := now.Format(item.layout)
		assert.Equal(t, date, timeutil.FormatBy(now, item.template))
	}

}

func TestFormatUnix(t *testing.T) {
	now := time.Now()
	want := now.Format("2006-01-02 15:04:05")

	assert.Equal(t, want, timeutil.FormatUnix(now.Unix()))
	assert.Equal(t, want, timeutil.FormatUnixBy(now.Unix(), timeutil.DefaultLayout))
}

func TestElapsedTime(t *testing.T) {
	nt := time.Now().Add(-time.Second * 3)
	num := timeutil.ElapsedTime(nt)
	i, err := strconv.ParseInt(num, 10, 64)
	assert.NoError(t, err)
	assert.Equal(t, 3000, i)
}
