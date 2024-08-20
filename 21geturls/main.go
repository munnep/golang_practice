package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("==========================")
	fmt.Println("testing with url")
	fmt.Println("==========================")

	// PerformGetRequest("http://localhost:8000/get")
	// PerformPostJsonRequest("http://localhost:8000/post")
	response := PerformPostFormRequest("http://localhost:8000/postform")

	fmt.Println(response, "coming from the PerformPostFormRequest")
	
}

func PerformPostFormRequest(myurl string) string {

	fmt.Println("==========================")
	fmt.Println("testing with Form POST")
	fmt.Println("==========================")
	// fake form data
    data := url.Values{}
		data.Add("firstname", "Patrick")
		data.Add("lastname", "test")
		data.Add("email", "Patrick@test.com")
		data.Add("email", "Patrick@test.com")
		data.Add("age", "41")
	

	response, err := http.PostForm(myurl, data)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(response.Body)
	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	response.Body.Close()
	return string(content)
}

func PerformPostJsonRequest(myurl string) {

	fmt.Println("==========================")
	fmt.Println("testing with Json POST")
	fmt.Println("==========================")
	// fake json payload
	requestBody := strings.NewReader(`
	  {
		"webserver":"patrick"
	  }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(response.Body)
	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	response.Body.Close()
}

func PerformGetRequest(myurl string) {

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	// option 1 for printing the content
	// body, _ := io.ReadAll(response.Body)
	// fmt.Println("Response from the server: ")
	// fmt.Println("Response code:", response.StatusCode)
	// fmt.Println(string(body))

	// option 2:
	var responseString strings.Builder
	body2, _ := io.ReadAll(response.Body)
	responseString.Write(body2)
	fmt.Println("using option 2:", responseString.String())

	defer response.Body.Close()
}
