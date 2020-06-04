package tool

import "time"

func NowDateTime() string {
    return DateTime(time.Now())
}

func NowDate() string {
    return Date(time.Now())
}

func NowTime() string {
    return Time(time.Now())
}

func DateTime(time time.Time) string {
    return time.Format("2006-01-02 15:04:05")
}

func Date(time time.Time) string {
    return time.Format("2006-01-02")
}

func Time(time time.Time) string {
    return time.Format("15:04:05")
}

func NowDateTimeWithLabel() (string, string) {
    return "date", DateTime(time.Now())
}

func ParseDateTime(timeStr string) time.Time {
    result, err := time.Parse("2006-01-02 15:04:05", timeStr)
    if nil != err {
        panic(err)
    }
    return result
}

func ParseDate(str string) time.Time {
    result, err := time.Parse("2006-01-02", str)
    if nil != err {
        panic(err)
    }
    return result
}

func ParseTime(str string) time.Time {
    result, err := time.Parse("15:04:05", str)
    if nil != err {
        panic(err)
    }
    return result
}
