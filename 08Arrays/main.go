package main

import "fmt"

func main() {
	fmt.Println("Exploring Arrays in Go")

	// array declaration
	var fleetPanda [10]string
	fleetPanda[0] = "Tech"
	fleetPanda[1] = "Marketing"
	fleetPanda[4] = "Products"

	fmt.Println("Team list is:", fleetPanda)
	fmt.Println("Length of array:", len(fleetPanda))

	var noOfTeams [5]int

	noOfTeams[0] = 5
	noOfTeams[1] = 10
	noOfTeams[4] = 7
	fmt.Println("No of teams is:", noOfTeams)
}
