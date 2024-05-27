package main

import "fmt"

func main() {
	fmt.Println("If ... Else If ... Else ...")
	age := 22
	var result string

	if age < 18 {
		result = "A Kid"
	} else if age > 18 {
		result = "An Adult"
	} else {
		result = "Exact 18"
	}

	fmt.Println(result)

	// Used for web request handling
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {}
}
