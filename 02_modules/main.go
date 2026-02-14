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

	//alias types
	type Age int
	var myAge Age = 25
	fmt.Printf("The value of myAge is: %d \n", myAge)
	fmt.Printf("The type of myAge is: %T \n", myAge)

	//here rune is an alias for int32 and byte is an alias for uint8
	//rune is used to represent a Unicode code point, while byte is used to represent a single byte of data
	var myRune rune = 'A'
	var myByte byte = 255
	fmt.Printf("The value of myRune is: %c \n", myRune)
	fmt.Printf("The type of myRune is: %T \n", myRune)
	fmt.Printf("The value of myByte is: %d \n", myByte)
	fmt.Printf("The type of myByte is: %T \n", myByte)
	//what is the difference between rune and byte in golang
	//rune is used to represent a Unicode code point, which can be a character from any language or symbol, while byte is used to represent a single byte of data, which can be a character from the ASCII character set or any other binary data. Runes are typically used for text processing, while bytes are used for low-level data manipulation.
	//what is the unicode code point
	//A Unicode code point is a unique number assigned to each character in the Unicode character set. It is used to represent characters from all writing systems, including letters, digits, symbols, and emojis. Each code point is represented as a hexadecimal value, and it can be encoded in different ways (such as UTF-8, UTF-16, etc.) for storage and transmission.
	//in simple terms, a Unicode code point is a way to represent characters from different languages and symbols in a standardized way, allowing for consistent encoding and decoding of text across different platforms and systems.

}
