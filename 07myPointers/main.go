package main

import "fmt"

func main() {
	fmt.Println("Exploring Pointers in Go....")

	// Declaring pointer * here the pointer stores integer value
	var ptr *int

	fmt.Println("What values does this pointer holds?", ptr)

	// & == reference
	myNumber := 45

	// reference myNumber to the pointer
	var pointingToMemoryAddress = &myNumber
	fmt.Println("Value of pointing Pointer is ", *pointingToMemoryAddress) // I want to see what's inside the pointer == value
	fmt.Println("Address of pointer is ", pointingToMemoryAddress)         // I want to see what is the memory address of the pointer

	*pointingToMemoryAddress = *pointingToMemoryAddress * 5

	fmt.Println("New value of myNumber is", myNumber)

}
