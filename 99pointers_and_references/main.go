package main

import "fmt"

func main() {
	x := "hello"
	y := &x

	fmt.Println("this is the variable x: ", x)
	// fmt.Println("this is the variable *x: ", *x) NOT possible
	fmt.Println("this is the reference y: ", y)
	fmt.Println("this is the pointer *y: ", *y)

	x = "world"
	fmt.Println("this is the variable x: ", x)
	fmt.Println("this is the pointer *y: ", *y)

	*y = "really?"
	fmt.Println("this is the variable x: ", x)
	fmt.Println("this is the pointer *y: ", *y)

	fmt.Println("======See the effect and usage of pointers ==================")
	var a string = "lotte"
	fmt.Println("before: ", a)
	alterString(a)
	fmt.Println("after: ", a)

	fmt.Println("****************************************")
	var b string = "lotte"
	fmt.Println("before: ", b)
	alterString2(&b)
	fmt.Println("after: ", b)

	// When using functions the pointers become important
	var g string = "hello"
	g = "world"
	fmt.Println(g)

	// versus
	var h string = "hello"
	h_pointer := &h
	*h_pointer = "world"
	fmt.Println(h)

}

func alterString(alterstring string) {
	alterstring = "Patrick"
	fmt.Println(alterstring)
}

func alterString2(alterstring *string) {
	*alterstring = "Patrick"
	fmt.Println(alterstring)
}
