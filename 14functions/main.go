package main

import "fmt"

func main() {
	fmt.Println("Functions!!")
	greet()

	result := adder(3, 5)
	fmt.Println("Result is", result)

	proResult, proResultMessage := proAdder(2, 5, 6, 7, 2, 4)
	fmt.Println(proResultMessage, proResult)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Pro Adder Result Successful"
}

func greet() {
	fmt.Println("Hello RIVO")
}
