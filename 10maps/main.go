package main

import "fmt"

func main() {
	fmt.Println("welcome to maps examples")

	// maps is key value
	languages := make(map[string]string)

	languages["RB"] = "Ruby"
	languages["JV"] = "Java script"
	languages["PY"] = "Python"

	fmt.Println("all languages:", languages)
	fmt.Println("The language for PY is ", languages["PY"])

	// delete value
	fmt.Println("delete the RB key/value")
	delete(languages, "RB")
	fmt.Println("all languages:", languages)

	// lets loop over this
	for key, value := range languages {
		fmt.Println("This is the key", key, "and the value", value)
		//or as follow
		fmt.Printf("This is the key %v and the value %v\n", key, value)
	}

	fmt.Println("==========================================")
	fmt.Println("slightly different loop example")
	fmt.Println("==========================================")
	// lets loop over this again but we don't care about the key
	for _, value := range languages {
		fmt.Println("This is the value", value)
		//or as follow
		fmt.Printf("This is the value %v\n", value)
	}
}
