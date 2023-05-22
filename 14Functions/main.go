package main

import "fmt"

func main() {
	fmt.Println("Exploring Functions..")

	result := adder(3, 4)
	result2 := multiplier(4, 5)

	fmt.Printf("Adder is %v \t and Multiplier is %v\n", result, result2)

	result3 := multipleValueTaker(3, 4, 5, 6, 7)
	fmt.Printf("Result from multipleValueTaker is %v \t", result3)
}

func adder(a int, b int) int {
	return a + b

}

func multiplier(x int, y int) int {
	return x * y
}

func multipleValueTaker(values ...int) int {

	totalCount := 0

	for _, value := range values {
		totalCount += value
	}
	return totalCount

}
