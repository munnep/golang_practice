package main

import (
	"fmt"
	"time"
)

func main() {

	current_time := time.Now()
	fmt.Println("Just mentioning the time:", current_time)

	// to print a nice time you always have to use the date
	// 01-02-2006 15:04:05 Monday     // Yes this is stupid instead of MM-DD-YYYY HH24:Mi:ss
	fmt.Println("Just mentioning the time:", current_time.Format("01-02-2006 15:04:05 Monday"))

	var createdDate time.Time = time.Date(2020, time.April, 12,23,23,0,0, time.UTC)
	fmt.Println(createdDate.Format("Monday"))
	fmt.Printf("type %T", createdDate)
}
