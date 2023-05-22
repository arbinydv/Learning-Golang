package main

import "fmt"

func main() {

	fmt.Println("******Exploring Basics of Defer******")
	fmt.Println("Hello ")           // this will get printed in normal fashion
	defer fmt.Println("Kubernetes") // first in last out
	defer fmt.Println("Golang")     // last in first out
}
