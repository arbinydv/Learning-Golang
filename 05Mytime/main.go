package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exploring Time pacakge in GO....")

	// present time

	presentTime := time.Now()

	fmt.Println("Current time: ", presentTime)

	// Formatting time based on user preferences
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	// Formatting Date time and Day as well
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//creating custom date and then formatting them to make it look cleaner

	customDate := time.Date(2023, time.December, 10, 20, 20, 20, 0, time.UTC)

	fmt.Println(customDate)

	//formatting it to look more clean

	fmt.Println(customDate.Format("01-02-2006 15:04:05 Monday"))

	// RubyDate formatting
	fmt.Println(presentTime.Format(time.RubyDate))
}
