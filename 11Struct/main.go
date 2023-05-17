package main

import "fmt"

// type Name of Struct and Type of Struct
// no inheritance in Go so no Parent Child concept exists in Go :(
type User struct {
	Name   string
	Age    int64
	Status bool
	Email  string

	// all these Fields are CAP so they can be accessed by anybody
}

func main() {
	fmt.Println("Starting with struct in Go...")

	// creating a new user using struct

	arbin := User{"Arbin", 26, true, "arbin@example.com"}

	// Default Printing
	fmt.Println("Details:", arbin)

	// detailed description using +
	fmt.Printf("Arbind Details are : %+v\n", arbin)

	// printing requried value only

	fmt.Printf("Name is %v\n Age is %v\n", arbin.Name, arbin.Age)

}
