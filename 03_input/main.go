package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to Go programming!"
	fmt.Println(welcomeMessage)

	reder := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	//comma ok syntax
	name, _ := reder.ReadString('\n')//it is actually like try catch in js
	fmt.Printf("Hello, %s! Nice to meet you.\n", name)
}
