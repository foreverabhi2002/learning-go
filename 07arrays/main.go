package main

import "fmt"

func main() {
	fmt.Println("Arrays!!")

	var reList [4]string

	reList[0] = "meteor"
	reList[1] = "classic"
	reList[2] = "interceptor"

	fmt.Println("RE List", reList)            // empty white space to denote the existence of space in array
	fmt.Println("RE List count", len(reList)) // to find length of the array

	// direct initialization at declaration
	var engineList = [5]string{"110cc", "125cc", "180cc", "200cc", "350cc"}
	fmt.Println("Engine List", engineList)
}
