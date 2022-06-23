package tool

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

const (
    dateTimeStr = "2006-01-02 15:04:05"
    dateStr     = "2006-01-02"
    timeStr     = "15:04:05"
)

func TestNowDateTime(t *testing.T) {
    dateTime := NowDateTime()
    assert.NotEmpty(t, dateTime)
    assert.Equal(t, 19, len(dateTime))
}

func TestNowDate(t *testing.T) {
    date := NowDate()
    assert.NotEmpty(t, date)
    assert.Equal(t, 10, len(date))
}

func TestNowTime(t *testing.T) {
    nowTime := NowTime()
    assert.NotEmpty(t, nowTime)
    assert.Equal(t, 8, len(nowTime))
}

func TestParseNowDateTime(t *testing.T) {
    dateTime := ParseDateTime(dateTimeStr)
    assert.NotEmpty(t, dateTime)
    date, month, day := dateTime.Date()
    clock, min, sec := dateTime.Clock()
    assert.Equal(t, date, 2006)
    assert.Equal(t, time.January, month)
    assert.Equal(t, 2, day)
    assert.Equal(t, 15, clock)
    assert.Equal(t, 4, min)
    assert.Equal(t, 5, sec)
}

func TestParseNowDate(t *testing.T) {
    d := ParseDate(dateStr)
    assert.NotEmpty(t, d)
    date, month, day := d.Date()
    assert.Equal(t, date, 2006)
    assert.Equal(t, time.January, month)
    assert.Equal(t, 2, day)
}

func TestParseNowTime(t *testing.T) {
    nowTime := ParseTime(timeStr)
    assert.NotEmpty(t, nowTime)
    clock, min, sec := nowTime.Clock()
    assert.Equal(t, 15, clock)
    assert.Equal(t, 4, min)
    assert.Equal(t, 5, sec)
}
