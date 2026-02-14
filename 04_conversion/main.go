package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza hunt!")
	fmt.Println("Please give the rating of the pizza from 1 to 10:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, you rated the pizza as:", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error parsing rating:", err)
	} else {
		fmt.Println("The numeric rating is:", numrating+1)
	}
}
