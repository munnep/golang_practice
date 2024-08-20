package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)

func main() {

	fmt.Println("working with web requests")

	res, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
 
    // website body is in bytes and you need to convert this to string
	// content := string(body)
	// fmt.Println(content)

	fmt.Println(string(body))

	defer res.Body.Close() //always have to do this last

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", body)
	fmt.Println("The status returned was:", res.Status)

	


}
