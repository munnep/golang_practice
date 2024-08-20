package main

import (
	"fmt"
	"strings"
)

// Methods operate on types. Mainly on a structure

// There are two types of receivers:

// Value Receiver: This means the method operates on a copy of the value, so the original value is not modified.
// Pointer Receiver: This means the method operates on a pointer to the value, allowing the method to modify the original value.

// Define a struct type.
type Rectangle struct {
	width  float64
	height float64
}

type Person struct {
	firstName string
	lastName  string
}

// Define a method with a value receiver.
func (r Rectangle) area() float64 {
	// The method can access the fields of the Rectangle type.
	return r.width * r.height
}

// Define a method with a pointer receiver.
func (p *Person) upper() {
	// The method modifies the fields of the original Person instance.
	p.firstName = strings.ToUpper(p.firstName)
}

func main() {
	// Create an instance of Rectangle.
	rect := Rectangle{width: 10, height: 5}

	// Call the area method on the rect instance.
	// area := rect.area()

	// Print the result.
	fmt.Println("Area:", rect.area()) // Output: Area: 50

	// Create an instance of person
	patrick := Person{firstName: "patrick", lastName: "test"}

	patrick.upper()
	fmt.Println("firstname: ", patrick.firstName)

}
