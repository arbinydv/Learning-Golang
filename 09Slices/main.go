package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Exploring slices...")

	// dynamic mememory allocation
	var pandaType = []string{"chinese", "nepali", "taiwanese", "filipinos"}

	fmt.Printf("Type of pandaType is %T\n", pandaType)

	// appending values inside slice
	pandaType = append(pandaType, "japanese", "korean")

	fmt.Println("Types of Pandas in Asia:", pandaType)

	// slicing method
	fmt.Println("Starting from Position{ inclusive} to Position {exclusive} : ", pandaType[1:5])

	highScore := make([]int, 4)
	highScore[0] = 100
	highScore[1] = 89
	highScore[2] = 99

	highScore = append(highScore, 95, 85)

	fmt.Println("What are values in highScore?", highScore)
	sort.Ints(highScore)
	fmt.Println("Sorted version of highScore", highScore)

	// removing value from slice

	var courses = []string{"reactjs", "python", "javascript"}
	fmt.Println("Printing all values:", courses)
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
