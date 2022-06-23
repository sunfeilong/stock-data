package tool

import "time"

var GMT = time.FixedZone("GMT", +8*60*60)

func NowDateTime() string {
    return DateTime(time.Now().In(GMT))
}

func NowDate() string {
    return Date(time.Now().In(GMT))
}

func NowTime() string {
    return Time(time.Now().In(GMT))
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
    return "date", DateTime(time.Now().In(GMT))
}

func ParseDateTime(timeStr string) time.Time {
    result, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, GMT)
    if nil != err {
        panic(err)
    }
    return result
}

func ParseDate(str string) time.Time {
    result, err := time.ParseInLocation("2006-01-02", str, GMT)
    if nil != err {
        panic(err)
    }
    return result
}

func ParseTime(str string) time.Time {
    result, err := time.ParseInLocation("15:04:05", str, GMT)

    if nil != err {
        panic(err)
    }
    return result
}
