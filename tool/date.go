package tool

import "time"

func Now() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func NowWithLabel() (string, string) {
    return "date", time.Now().Format("2006-01-02 15:04:05")
}
