package main

import (
	"fmt"
	"time"
	"concurrency/hello"
)

func main() {
   go greeting("hello")
   go greeting("world")

   time.Sleep(3 * time.Second)


   // using a different module 
   hello.Say("worked using module")

}

func greeting(s string)  {
	for i := 0; i < 6; i++ {
		fmt.Println("Message is: ", s)
		time.Sleep(3 * time.Millisecond)
	}
}