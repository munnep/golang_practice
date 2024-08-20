package main

import (
	"fmt"
	"io"
	"os"
)

const fileName_str string = "./myfile.txt"

func main() {

	fmt.Println("working with files")

	content := "this is data in the file - patrick\n"
	// create a file
	file, err := os.Create(fileName_str)
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is:", length)
    
	// extra line with no details or error handling
	io.WriteString(file, "extra line\n")

	defer file.Close()

	// read the file again using a function
	readfile(fileName_str)
}

func readfile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("===========================")
	fmt.Println("Reading data from the file")
	fmt.Println("===========================")
	fmt.Printf("File contents: %s", data)
}
