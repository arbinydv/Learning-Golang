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
	fmt.Println("Starting with method in Go...")

	// creating a new user using struct

	arbin := User{"Arbin", 26, true, "arbin@example.com"}
	fmt.Println("Details:", arbin)

	//calling method
	arbin.GetAge()
	arbin.GetEmail()

}

// creating a method
func (u User) GetAge() {

	result := "No"

	if u.Age >= 18 {
		result = "Yes"
	}
	fmt.Println("Is user above drinking age? ", result)
}

// method 2
// user.Email is passed as value here { pointer concept so it doesn't edit the actual User.Email which is arbin@....}
func (u User) GetEmail() {
	u.Email = "havann@example.com"

	fmt.Println("Email is:", u.Email)
}
