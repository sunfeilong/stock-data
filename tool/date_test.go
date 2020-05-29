package tool

import (
    "github.com/stretchr/testify/assert"
    "testing"
    time2 "time"
)

const (
    dateTimeStr = "2006-01-02 15:04:05"
    dateStr     = "2006-01-02"
    timeStr     = "15:04:05"
)

func TestNowDateTime(t *testing.T) {
    time := NowDateTime()
    assert.NotEmpty(t, time)
    assert.Equal(t, 19, len(time))
}

func TestNowDate(t *testing.T) {
    time := NowDate()
    assert.NotEmpty(t, time)
    assert.Equal(t, 10, len(time))
}

func TestNowTime(t *testing.T) {
    time := NowTime()
    assert.NotEmpty(t, time)
    assert.Equal(t, 8, len(time))
}

func TestParseNowDateTime(t *testing.T) {
    time := ParseDateTime(dateTimeStr)
    assert.NotEmpty(t, time)
    date, month, day := time.Date()
    clock, min, sec := time.Clock()
    assert.Equal(t, date, 2006)
    assert.Equal(t, time2.January, month)
    assert.Equal(t, 2, day)
    assert.Equal(t, 15, clock)
    assert.Equal(t, 4, min)
    assert.Equal(t, 5, sec)
}

func TestParseNowDate(t *testing.T) {
    time := ParseDate(dateStr)
    assert.NotEmpty(t, time)
    date, month, day := time.Date()
    assert.Equal(t, date, 2006)
    assert.Equal(t, time2.January, month)
    assert.Equal(t, 2, day)
}

func TestParseNowTime(t *testing.T) {
    time := ParseTime(timeStr)
    assert.NotEmpty(t, time)
    clock, min, sec := time.Clock()
    assert.Equal(t, 15, clock)
    assert.Equal(t, 4, min)
    assert.Equal(t, 5, sec)
}
