package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	endpoint := "http://localhost:8080/"
	// Request data container
	data := url.Values{}
	// Command line prompt
	fmt.Println("Enter full URL")
	// Opens stream reading from the prompt
	reader := bufio.NewReader(os.Stdin)
	// Reading line from prompt
	long, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	long = strings.TrimSuffix(long, "\n")
	// Filling container with data
	data.Set("url", long)
	// Add http client
	client := &http.Client{}
	// Writing request.
	// POST request has to contain both headers and body.
	// Body has to be a source for io.Reader.
	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	// Defines encoding in the request's header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Sending request and getting response
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	// Print response status code
	fmt.Println("Status code:", response.Status)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	// Reading the stream from the body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// Printing the stream
	fmt.Println(string(body))
}
