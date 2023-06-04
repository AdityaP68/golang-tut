package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// this thing is fixed -- 01-02-2006 15:04:05 Monday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10,23,23,0,0,time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	 
}
