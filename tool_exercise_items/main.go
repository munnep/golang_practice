package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "slices"
)

// use constants as it is more easy to rename later in the future

const (
	DATA_FILE     string = "items"
	ACTION_ADD    string = "add"
	ACTION_LIST   string = "list"
	ACTION_REMOVE string = "remove"
)

func main() {
	if len(os.Args) < 2 {
		exit("No correct argument given")
	}
	// fmt.Println("test")

	// You can use the os.Args package to pass arguments
	// this is a list of values where the index 1 is the first argument given
	// os.Args
	// fmt.Println(os.Args[1])

	action := os.Args[1]
	// fmt.Println(action)

	switch action {
	case ACTION_ADD:
		if len(os.Args) < 3 {
			exit("Need at least 1 item to add")
		}

		items := os.Args[2:]

		add(items)
	case ACTION_LIST:
		list()
	case ACTION_REMOVE:
		if len(os.Args) < 3 {
			exit("Need at least 1 item to remove")
		}

		items := os.Args[2:]

		remove(items)
	default:
		fmt.Println("Valid arguments are: ", ACTION_ADD, ACTION_LIST, ACTION_REMOVE)
	}
}

func remove(items []string) {
	fmt.Println("removing items", items)

	content, err := os.ReadFile(DATA_FILE)
	check(err)
	// trim for removing the last empty lines
	// split to have every lines separate
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	// fmt.Println("print lines", lines)

	// create a map for the items
	remove := make(map[string]bool)

	for _, item := range items {
		remove[item] = true
	}

	file, err := os.OpenFile(DATA_FILE, os.O_CREATE|os.O_WRONLY, 0664)
	check(err)
	defer file.Close()

	// empty the file
	file.Truncate(0)

	for _, item := range lines {
		if !remove[item] {
			file.WriteString(item + "\n")
		}
	}
}

// fmt.Println(remove)

func add(items []string) {
	fmt.Println("items to add:", items)

	file, err := os.OpenFile(DATA_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	check(err)
	defer file.Close()

	// make sure that when something is done it is at the end of the file
	// not need if you use the os.O_APPEND from the previous step
	// file.Seek(0,2)

	// OPTION 1
	// for _, item := range items {
	// 	_, err := file.WriteString(item + "\n")
	// 	check(err)
	// }

	// OPTION 2 for high volume writes
	writer := bufio.NewWriter(file)
	for _, item := range items {
		_, err := writer.WriteString(item + "\n")
		check(err)
	}
	err = writer.Flush()
	check(err)

}

func list() {
	// option 1: below is an easy option of opening the file
	// file, err := os.Open("items")

	// option 2: this is a possible option with more control
	// file, err := os.OpenFile("items", os.O_CREATE|os.O_RDONLY, 0664)
	// check(err)
	// // fmt.Println(file)
	// cleanup
	// defer file.Close()

	// info, err := file.Stat()
	// check(err)
	// fmt.Println("Size of the file:", info.Size())

	// content := make([]byte, info.Size())
	// count, err := file.Read(content)
	// check(err)
	// fmt.Println(count)
	// fmt.Println(string(content))

	// Option 3: using the readfile option which is more simple, but it will read the entire file into memory. (imagine a file of 4GB to read???)

	content, err := os.ReadFile(DATA_FILE)
	check(err)
	fmt.Println(strings.TrimSpace(string(content)))

	// option 4: which can be used with large files to read one by one
	// file, err := os.OpenFile(DATA_FILE, os.O_CREATE|os.O_RDONLY, 0664)
	// check(err)
	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	// scanner.Split(bufio.ScanLines)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// check(scanner.Err())

}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
