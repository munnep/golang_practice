package main

import (
	"fmt"
	// "log"
	"os"
	// "strconv"
	"encoding/json"
)

func main() {
	fmt.Println("test with bytes and conversions")

	// var number int = 86
	// var letter string = "T"

	// fmt.Println(number)

	// fmt.Println([]byte(letter))
	// fmt.Println(string(number))
	// fmt.Println(strconv.Itoa(number))

	readJSONfile("example.json")

}

func readJSONfile(filename string) {
	body, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("print data type is %T", data)
	// fmt.Println(string(data))

	// parse the JSON file
	var test map[string]interface{}
	json.Unmarshal(body, &test)

	// fmt.Println(test)


	// Extract questions from the quiz data
	quiz := test["quiz"].(map[string]interface{})
	extractQuestions(quiz)

}

// Function to extract questions from a given map
func extractQuestions(quizMap map[string]interface{}) {
	for _, category := range quizMap {
		categoryMap := category.(map[string]interface{})
		for _, questionData := range categoryMap {
			questionMap := questionData.(map[string]interface{})
			question := questionMap["question"].(string)
			fmt.Println("Question:", question)
		}
	}
}
