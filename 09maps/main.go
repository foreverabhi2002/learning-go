package main

import "fmt"

func main() {
	fmt.Println("Maps!!")

	languages := make(map[string]string) // map declaration

	languages["js"] = "Javascript" // initialization
	languages["ts"] = "Typescript"
	languages["py"] = "Python"

	fmt.Printf("languages: %v\n", languages)
	fmt.Printf("language %v\n", languages["js"])

	delete(languages, "py") // delete from array
	fmt.Printf("updated languages: %v\n", languages)

	for key, value := range languages {
		fmt.Printf("key is %v, and the value is %v\n", key, value)
	}

}
