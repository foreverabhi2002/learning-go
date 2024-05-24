package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Rivo")

	fmt.Println("Enter your year of birth")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // by default the user input type is of string

	fmt.Println("YOB is", input)

	numericYob, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Your age is: ", 2024-numericYob)
	}
}
