package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Staring with type conversion...")
	fmt.Println("Please enter a rating between  0 and 10")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your response", input)

	// input will be of type string and converting it into float64

	toFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Coverted your rating into float", toFloat)
	}

}
