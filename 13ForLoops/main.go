package main

import "fmt"

func main() {
	fmt.Println("Exploring Forloops in Go..")

	myFavouriteDays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri"}
	fmt.Println("myFavouriteDays = ", myFavouriteDays)

	// Normal looping

	for d := 0; d <= len(myFavouriteDays)-1; d++ {
		fmt.Println(myFavouriteDays[d])
	}

	// using range

	for index := range myFavouriteDays {
		fmt.Println(myFavouriteDays[index])
	}

	// using index, value with range

	for index, days := range myFavouriteDays {
		fmt.Printf("For index %v The day is %v\n", index, days)
	}

}
