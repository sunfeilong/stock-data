package main

import (
	"fmt"
)

func main() {

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
}
