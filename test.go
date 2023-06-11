package main

import (
	"fmt"
	"time"
)

func main() {
	現在時間Sstr := "2023-05-26 00:00:00"
	timeTemplate := "2006-01-02 15:04:05"
	deadline, _ := time.ParseInLocation(timeTemplate, deadlineString, time.Local)
	cnLocal, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(cnLocal)
	if now.Before(deadline) {
		fmt.Println("1212")
	}

}
