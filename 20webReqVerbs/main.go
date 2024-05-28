package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web req verbs")
	// GetRequest("http://localhost:3000")
	// PostJsonRequest("http://localhost:3000/add")
	PostFormRequest("http://localhost:3000/add")
}

func GetRequest(myUrl string) {
	res, err := http.Get(myUrl)
	checkForError(err)
	defer res.Body.Close()
	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length:", res.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))
	fmt.Println("Bytecount:", byteCount)
	fmt.Println("content", responseString.String())
}

func PostJsonRequest(myUrl string) {
	requestBody := strings.NewReader(`
	{
		"title": "Hello",
		"description": "Namaste"
	}
	`)

	res, err := http.Post(myUrl, "application/json", requestBody)
	checkForError(err)

	defer res.Body.Close()
	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	responseString.Write(content)
	fmt.Println(responseString.String())
}

func PostFormRequest(myUrl string) {
	data := url.Values{}
	data.Add("title", "foreverabhi")
	data.Add("description", "description of foreverabhi and rivolabs")

	res, err := http.PostForm(myUrl, data)
	checkForError(err)
	defer res.Body.Close()

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	responseString.Write(content)
	// fmt.Println(string(content))
	fmt.Println(responseString.String())
}

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}
