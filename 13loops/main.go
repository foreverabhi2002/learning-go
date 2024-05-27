package main

import "fmt"

func main() {
	fmt.Println("Loooooooooops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println("days", days)

	// Basic looping technique
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("day", days[d])
	// }

	// Loops with range
	// for day := range days {
	// 	fmt.Println(days[day])
	// }

	// Index-value loop using range
	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// Golang's Alternative to while loop
	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto rivo
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is", rogueValue)
		rogueValue++ //no pre increment or decrement in values
	}

rivo:
	fmt.Println("Printing RIVO")

}
