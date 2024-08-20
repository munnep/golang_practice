package main

import (
	"fmt"
	"net/url"
)

const website string = "https://google.com:443/search?q=golang"

func main() {
	fmt.Println("working with urls example:", website)

	result_url, _ := url.Parse(website)

	fmt.Println(result_url.Scheme)
	fmt.Println(result_url.Host)
	fmt.Println(result_url.Port())
	fmt.Println(result_url.Path)
	fmt.Println(result_url.RawQuery)

	// fmt.Println(result_url.Query())

	qparameters := result_url.Query()

	// print the query part
	fmt.Println(qparameters["q"])

	for key, value := range qparameters {
		fmt.Println("the key: ", key)
		fmt.Println("the value:", value)
	}

	// you want to reference the url

	anotherUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/search",
		RawQuery: "q=patrick",
	}

	// print it out
	fmt.Println("notice your miss the rawpath:", anotherUrl)
	// you will need to convert it to string to avoud issues
	// anotherUrl2 := anotherUrl.String()
	fmt.Println("Now it is complete url", anotherUrl.String())
}
