package main

import (
    "fmt"
    "log"
    "os"
)

var logger *log.Logger

func init() {
    logger = log.New(os.Stdout, "[main] ", log.Llongfile|log.LstdFlags)
}

func main() {
    logger.Println("项目启动")

    fmt.Println("start ...")
    first := make(chan bool, 1)
    second := make(chan bool, 1)
    third := make(chan bool, 1)

    go func() {
        fmt.Println("定时抓取股票列表...")
        first <- true
    }()

    go func() {
        <-first
        fmt.Println("抓取每天数据...")
        second <- true
    }()

    go func() {
        <-second
        fmt.Println("抓取实时数据...")
        third <- true
    }()

    <-third
    fmt.Println("end ...")

    logger.Println("项目运行结束")
}
