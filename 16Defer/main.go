package main

import "fmt"

func main() {

	fmt.Println("******Exploring Basics of Defer******")
	fmt.Println("Hello ")           // this will get printed in normal fashion
	defer fmt.Println("Kubernetes") // first in last out
	defer fmt.Println("Golang")     // last in first out

	// reverse print showing defer example in use

	values := []int{1, 2, 3, 4, 5}

	for i, _ := range values {
		defer fmt.Printf("Number is: %v\n", i)
	}

	// Defer works in LIFO order
}
