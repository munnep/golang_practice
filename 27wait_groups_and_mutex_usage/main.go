package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals []string
var wg sync.WaitGroup
var mut sync.Mutex

func main() {

	websites := []string{
		"https://google.com",
		"https://hashicorp.com",
		"https://microsoft.com",
		"https://go.dev",
		"https://github.com",
		"https://example.com",
	}

	for _, web := range websites {
		// fmt.Println("Testing website: ", web)
		wg.Add(1)
		go getStatusCode(web, &wg, &mut)
	}
	// wait for all go routines to finish
	wg.Wait()

	fmt.Println(signals)

}

func getStatusCode(endpoint string, wg *sync.WaitGroup, mut *sync.Mutex) {
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Something wrong:", err)
	} else {
		//using %d because we expect an integer
		fmt.Printf("%d statuscode for website %s\n", result.StatusCode, endpoint)

		// this is needed with the go routines otherwise we have to potential of different go routines trying to alter the same piece of memory
		mut.Lock()
        signals = append(signals, endpoint)
		mut.Unlock()
	}
    // mark the go routine as done
	defer wg.Done()

}
