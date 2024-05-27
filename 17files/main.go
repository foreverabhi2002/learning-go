package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in GoLang")
	content := "This needs to go in a file - rivolabs.in"

	file, err := os.Create("./rivo.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("Length is", length)
	defer file.Close()
	readData := readFile("./rivo.txt")
	fmt.Println("read data\n", readData)
}

func readFile(filename string) string {
	databyte, err := os.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	return string(databyte)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
