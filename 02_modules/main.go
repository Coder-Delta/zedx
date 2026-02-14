package main

import "fmt"

func main() {
	var message string = "Hello, World!"
	fmt.Println(message)
	fmt.Printf("This is a simple Go program: %T \n", message)

	var isLoggedIn bool = true
	fmt.Printf("Is the user logged in? %t \n", isLoggedIn)
	fmt.Printf("The type of isLoggedIn is: %T \n", isLoggedIn)

	var string uint8 = 100 // range of uint8 is 0 to 255
	fmt.Printf("The value of string is: %d \n", string)
	fmt.Printf("The type of string is: %T \n", string)

	//has many types of data types in golang uint8, uint16, uint32, uint64 and int8, int16, int32, int64 unsigned and signed integers respectively 

	var smallFloat float32 = 3.141592653589793 // float32 has a precision of about 7 decimal places, so it will not be able to represent the full value of pi accurately
	fmt.Printf("The value of smallFloat is: %f \n", smallFloat)
	fmt.Printf("The type of smallFloat is: %T \n", smallFloat)

	var smallFloat64 float64 = 3.141592653589793 // float64 has a precision of about 15 decimal places, so it can represent the value of pi more accurately than float32
	fmt.Printf("The value of smallFloat64 is: %f \n", smallFloat64)
	fmt.Printf("The type of smallFloat64 is: %T \n", smallFloat64)

	//default and alias types
	var age int = 30
	fmt.Printf("The value of age is: %d \n", age)
	fmt.Printf("The type of age is: %T \n", age)

	//implicit type inference
	var name = "Alice"
	fmt.Printf("The value of name is: %s \n", name)
	fmt.Printf("The type of name is: %T \n", name)

	//short variable declaration
	country := "USA"
	fmt.Printf("The value of country is: %s \n", country)
	fmt.Printf("The type of country is: %T \n", country)
	// this decaration don't allow u to write the code in the global scope but others are allowed to write in the global scope
}
