package main

import "fmt"

func main() {
	fmt.Println("It's all about pointers")

	ptr := 2002

	var address = &ptr

	fmt.Println(ptr, address, *address) // 2002 0xc00000a0c8 2002
}
