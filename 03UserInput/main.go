package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	initialText := "Welcome to the console."
	fmt.Println(initialText)

	// Get userInput

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" What is your name? ")

	// comma OK   or error Ok syntax

	userInput, _ := reader.ReadString('\n')
	fmt.Println("The user Input is: ", userInput)

}
