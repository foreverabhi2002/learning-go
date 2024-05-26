package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("All about Slices")

	var reList = []string{"meteor", "classic", "interceptor", "continental"}

	fmt.Println("RE List", reList)

	reList = append(reList, "standard", "hunter")

	fmt.Println("Updated RE List", reList)

	// reList = append(reList[1:4]) // slices out before the 1st index and non inclusive 4th index and rest

	// fmt.Printf("After slicing RE List: %v\n", reList)

	highScores := make([]int, 4)

	highScores[0] = 10
	highScores[1] = 20
	highScores[2] = 25
	highScores[3] = 20
	highScores = append(highScores, 100, 45, 30) // can append to fixed slices length

	fmt.Printf("highScores: %v\n", highScores)

	sort.Ints(highScores)          // sorts integers from arrays
	sort.IntsAreSorted(highScores) // returns true if sorted
	fmt.Printf("highScores sorted: %v\n", highScores)

	// remove value from slices
	toRemoveIndex := 3
	reList = append(reList[:toRemoveIndex], reList[toRemoveIndex+1:]...) // to remove index from slice we need to slice the slice from both the sides
	fmt.Printf("reList: %v\n", reList)
}
