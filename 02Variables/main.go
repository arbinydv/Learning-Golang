package main

import "fmt"

// Global Scope constant declaration
// First Character is  alwaysUPPER CASE
const UserOTP string = "abdecf123"

func main() {
	// String Type
	var st string = "This is a string variable"

	// To find type of a variable
	fmt.Printf("The variable is of type:  %T \n", st)

	// Boolean
	var isGolang bool = true
	fmt.Printf("The variable is of type:  %T \n", isGolang)

	// Float32 and Float64 based on the precision || no of digits after the decimal point
	var smallFloat float32 = 344.56777667777
	fmt.Printf("The float value is: %f \n", smallFloat)
	fmt.Printf("The variable is of type: %T \n", smallFloat)

	// implict type of variable declaration LEXER automatically assign the variable data type base on the value user put

	var checker = "This is string "
	fmt.Println("implict type declaration:", checker)
	// checker = 5.6  cannot be assigned as LEXER treats checker variable as string

	// no var style declaration == declaration + initialization  varless operation

	noOfDaysLeftForDashain := 23
	fmt.Println("How many days to go for Dashain:", noOfDaysLeftForDashain)

	// NOTE : we're only allowed to use varless operation inside a function or method and it DOESN'T support global scope declaration

	fmt.Println("This is a login OTP", UserOTP)

}
