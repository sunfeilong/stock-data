package main

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "github.com/xiaotian/stock/pkg/tool"
    "testing"
    "time"
)

func TestNextRunDurationZh(t *testing.T) {
    ti := time.Now()
    times := 60 * 24 * 10000

    for i := 0; i < times; i++ {
        zh := nextRunDurationZh(ti)
        nextTime := ti.Add(zh)

        assert.NotEqual(t, time.Sunday, nextTime.Weekday())
        assert.NotEqual(t, time.Saturday, nextTime.Weekday())
        assert.Equal(t, 16, nextTime.Hour())
        assert.Equal(t, 30, nextTime.Minute())

        ti = ti.Add(time.Minute)
        if i%(60*24) == 0 {
            fmt.Println("当前进度: ", 1.0*i*100/times, " %")
        }
    }
}

func TestSpecial(t *testing.T) {

    dateTime := tool.ParseDateTime("2020-06-07 05:00:16")
    durationZh := nextRunDurationZh(dateTime)
    next := dateTime.Add(durationZh)
    assert.Equal(t, "2020-06-08 16:30:00", tool.DateTime(next))

    dateTime = tool.ParseDateTime("2020-06-05 04:09:46")
    durationZh = nextRunDurationZh(dateTime)
    next = dateTime.Add(durationZh)
    assert.Equal(t, "2020-06-05 16:30:00", tool.DateTime(next))

    dateTime = tool.ParseDateTime("2020-08-12 08:34:21")
    durationZh = nextRunDurationZh(dateTime)
    next = dateTime.Add(durationZh)
    assert.Equal(t, "2020-08-13 16:30:00", tool.DateTime(next))

}
