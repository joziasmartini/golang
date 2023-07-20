package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Making a GET request
	getRequest()

	// Making a POST request
	postRequest()
}

func getRequest() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Perform the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	// Close the response body when the function returns
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response
	fmt.Println("GET Response:")
	fmt.Println(string(body))
}

func postRequest() {
	url := "https://jsonplaceholder.typicode.com/posts"
	payload := []byte(`{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`)

	// Perform the POST request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}

	// Close the response body when the function returns
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response
	fmt.Println("POST Response:")
	fmt.Println(string(body))
}
