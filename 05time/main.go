package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now about time")

	presentTime := time.Now()

	fmt.Println("The time is", presentTime)

	// Formatting method
	fmt.Println("Format time", presentTime.Format("02-01-2006 Mon"))

	// Create another date
	createdDate := time.Date(2002, time.February, 10, 0, 0, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}
