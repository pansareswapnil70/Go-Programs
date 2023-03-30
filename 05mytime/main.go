package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang Time Assignment")
	currentTime := time.Now()
	fmt.Println("Current time is : ", currentTime)
	fmt.Println("Formatted Date is: ", currentTime.Format("02-01-2006 15:04:05 Monday"))
	Date := time.Date(2023, time.March, 23, 12, 23, 00, 00, time.UTC)
	fmt.Println("Created Date is: ", Date)
}
