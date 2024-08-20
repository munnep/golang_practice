package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "tomato", "peach"}

	fruitList = append(fruitList, "mango")

	fmt.Println(fruitList)

	// remove the index - value.....apple
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	test := &fruitList
	fmt.Println("pointer:", *test)

	highscores := []float64{23, 23423, 3}
	fmt.Println(highscores)

	highscores = append(highscores, 3, 221, 444)
	fmt.Println(highscores)

	sort.Float64s(highscores)
	fmt.Println(highscores)

	var courses = []string{"postgresql", "oracle", "sqlserver"}
	fmt.Println("output courses:", courses)
    fmt.Println("remove oracle from the list")
	// removing index 1
	index := 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("without oracle:", courses)
}
