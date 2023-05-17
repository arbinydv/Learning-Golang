package main

import "fmt"

func main() {

	fmt.Println("Starting with Maps in Go..")

	// create map using make

	languages := make(map[string]string)
	// Key Value addition in map
	languages["Js"] = "javascript"
	languages["Rb"] = "ruby"
	languages["C"] = "C++"

	fmt.Println("List of all languages:", languages)
	fmt.Println("Rb is :", languages["Rb"])

	// Deletion in Map
	delete(languages, "Rb")

	// loop through Maps
	//we can use comma, Ok syntax based on what we want from data
	for key, value := range languages {

		fmt.Printf(" Key is %s and value is %s\n", key, value)
	}

}
