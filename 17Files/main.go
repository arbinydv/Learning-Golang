package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Exploring file handling in Go..")
	content := "This is Happy. Happy likes Bananas. Happy plays hockey and basketball. Happy makes other happy. Happy is a good person. Be like Happy :)"

	// file creation

	file, err := os.Create("./happyBio.txt")
	if err != nil {
		panic(err) // This will abort further processing
	}

	// write content inside file
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("The content has been copied and file length is:", length)

	// reading file
	readFile("./happyBio.txt")

	defer file.Close() // better to use defer so that we can close once everything is done

}

// read file

func readFile(fileName string) {

	dataByte, err := ioutil.ReadFile(fileName) // everything is converted into Byte before they're typecasted

	if err != nil {
		panic(err)
	}
	fmt.Println("Raw data in dataBytes is ", dataByte)

	//coverting it back to string
	originalContent := string(dataByte)
	fmt.Println("\nReading content of the file:", originalContent)

}
