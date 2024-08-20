package main

import "fmt"

func main() {
	fmt.Println("Using functions")

	greeting()
	fmt.Println(greeting2("patrick"))

	adding_result := adding(5, 5)
	fmt.Println(adding_result)

	adding_complex_result := adding_complex(3,4,1,5,6,3,1,3)
	fmt.Println(adding_complex_result)
}

func adding(value1 int, value2 int) int {
	return value1 + value2
}

func adding_complex(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
    return total
}

func greeting() {
	fmt.Println("Hello and good morning")
}

func greeting2(name string) (string, string) {
	return "hello", name
}
