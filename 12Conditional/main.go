package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("\n  I: ifelse Condtion \t s S: switchCase Condtion")
	fmt.Printf("\nEnter the desired operation you want to explore.\n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//UpperCase Conversion
	userChoice := strings.ToTitle(input)

	// Here we are removing any trailing whitespaces from the input before comparison
	if strings.TrimRight(userChoice, "\n") == "I" {
		expoloreIfElseCondition()
	} else if strings.TrimRight(userChoice, "\n") == "S" {
		exploreSwitchCase()
	} else {
		fmt.Println("Oops! You entered something else :(")
	}

}
func expoloreIfElseCondition() {
	fmt.Println("*************************Exploring If Else condition in Go..***********************")

	fmt.Println("Enter a number to check the Timebomb status")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	counter, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 32)

	if counter < 10 {
		fmt.Println("Beaware of timebomb zone.......")
	} else if counter > 10 {
		fmt.Println("Oh You're safe and out of timebomb zone....:)")
	} else {
		fmt.Println("Run Run Run .......")
	}

}

func exploreSwitchCase() {
	fmt.Println("*************************Exploring Switch case in Go..***********************")
	// Generate random number
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of a dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can open the gate.")
	case 2:
		fmt.Println("Dice value is 2 and you can move to 2 spot forward")
	case 3:
		fmt.Println("Dice value is 3 and you can move to 3 spot forward")
	case 4:
		fmt.Println("Dice value is 4 and you can move to 4 spot forward")
	case 5:
		fmt.Println("Dice value is 5 and you can move to 5 spot forward")
	case 6:
		fmt.Println("Dice value is 6 and you got a chance to roll it again")
	default:
		fmt.Println("Pass the dice to next team mate")
	}

}
