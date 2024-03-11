package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2023, time.September, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local()) // convert UTC to local time

	formatter := "2006-01-02 15:04:05"
	value := "2002-09-12 14:18:37"
	valueTime, err := time.Parse(formatter, value)
	if err == nil {
		fmt.Println(valueTime)
	} else {
		fmt.Println(err.Error())
	}

	// get the specific of time
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Minute())
	fmt.Println(valueTime.Second())
}
