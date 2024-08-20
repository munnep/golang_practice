package main

import "fmt"

func main() {

	fmt.Println("testing with loops")

	days_slice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days_slice)

	days_map := make(map[string]string)

	days_map["Monday"] = "Crap.....start working"
	days_map["Tuesday"] = "Finshed one day"
	days_map["Wednesday"] = "Half way there"
	days_map["Thursday"] = "finishing things up"
	days_map["Friday"] = "dreaming about the weekend"

	// for i := 0; i < len(days_slice); i++ {
	// 	fmt.Println(days_slice[i])
	// }

	// for i := range days_slice {
	// 	fmt.Println(days_slice[i])
	// }

	// for k, v := range days_slice {
	// 	fmt.Printf("What I think of day %v is \"%v\" \n", k, v)
	// }

	just_a_number := 1

	for just_a_number < 10 {

		if just_a_number == 5 {
			just_a_number++ // we add +1 to make sure we go to 6
			fmt.Println("We hate this number. Not showing")
            continue
		}

		if just_a_number == 6 {
			fmt.Println("we jump using goto")
            goto label_example
		}
	

		fmt.Println(just_a_number)

		if just_a_number == 8 {
			fmt.Println("we cancel at number 8")
			break
		}
		just_a_number++
	}

label_example:
   fmt.Println("this is use of a label and goto")	
	// for k, v := range days_map {
	// 	fmt.Printf("What I think of day %v is \"%v\" \n", k, v)
	// }
}
