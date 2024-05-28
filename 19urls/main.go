package main

import (
	"fmt"
	"net/url"
)

const myUrl = "http://localhost:5525/docs?username=foreverabhi&&password=abhi2002"

func main() {
	fmt.Println("Handling urls in golang")
	fmt.Println("url", myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params: %T\n", qParams)

	fmt.Println(qParams["username"])

	for index, val := range qParams {
		fmt.Printf("Index id: %v - Param is: %v\n", index, val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "rivolabs.in",
		Path:    "/foreverabhi",
		RawPath: "user=admin",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
