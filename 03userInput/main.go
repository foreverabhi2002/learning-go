package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Rivo"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any thing: ")

	// comma ok || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for entering, ", input)
	fmt.Printf("Type of %T", input)
}
