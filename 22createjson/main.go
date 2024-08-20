package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// this is a normal structure but with a third paramater that json will take into account
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // I don't want to output this value
	Tags     []string `json:"tags,omitempty"` // if nil then don;t show at all
}

func main() {
	fmt.Println("==========================")
	fmt.Println("testing with JSON creation")
	fmt.Println("==========================")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "LearnCodeOnline.in", "csdf234", []string{"google", "js"}},
		{"Python Bootcamp", 299, "LearnCodeOnline.in", "abcwer234", nil},
	}

	fmt.Println("==========================")
	fmt.Println("Print as jsonencode")
	fmt.Println("==========================")
	// package this data as json
	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(finalJson))

	//    or bit nicer format
	// package this data as json
	fmt.Println("==========================")
	fmt.Println("print as json pretty")
	fmt.Println("==========================")
	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t") // with tab if you want "  " for spaces
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(finalJson2))

}

func DecodeJson() {
	fmt.Println("==========================")
	fmt.Println("Print as JSON DECODE")
	fmt.Println("==========================")

	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcocourse course

	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("valid json")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
		fmt.Println(lcocourse)
	} else {
		fmt.Println("not valid json")
	}

	// test, _ := json.MarshalIndent(lcocourse, "", "\t")
	// fmt.Println(string(test))

	// cases where you just want to store key value data from json
	fmt.Println("WORKING WITH KEY VALUE DATA")
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("The key: %v the value: %v \n", key, value)
	}

}
