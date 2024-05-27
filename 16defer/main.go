package main

import "fmt"

func main() {
	// defer process is in lifo structure
	defer fmt.Println("Whyyy?")    // 2
	fmt.Println("Deferrrrrrrrr!!") // 1
}
