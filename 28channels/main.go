package main

import (
	"fmt"
	"sync"
)

func main()  {

	fmt.Println("==========================")
	fmt.Println("working with channels")
	fmt.Println("==========================")
	
	wg := &sync.WaitGroup{}
	mychannel := make(chan int, 2)

	wg.Add(2)

	// receive only
	go func (ch <-chan int, wg *sync.WaitGroup)  {
		value, isChannelOpen := <-mychannel


		fmt.Println(value)
		fmt.Println(isChannelOpen)
		fmt.Println(<-mychannel)
		// fmt.Println(<-mychannel)
		wg.Done()
	}(mychannel, wg)
 
	// Send only channel
	go func (ch chan<- int, wg *sync.WaitGroup)  {
		
		mychannel <- 5 
		close(mychannel)
		// mychannel <- 6
		wg.Done()
	}(mychannel, wg)



	wg.Wait()

}