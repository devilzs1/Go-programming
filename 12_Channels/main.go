package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Learning about Channels & Deadlock using Go Programming")

	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// mut := &sync.Mutex{}

	// // Channel allows to pass the value when someone is listening to it
	// myChan <- 7
	// fmt.Println(<-myChan) // fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan send]

	// // Still you will face the same error
	// fmt.Println(<-myChan) // fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan send]
	// myChan <- 7



	// wg.Add(2)
	// go func(ch chan int, wg *sync.WaitGroup){
		
	// 	fmt.Println("My channel value : ", <-myChan)
	// 	wg.Done()
	// }(myChan, wg)
	// go func(ch chan int, wg *sync.WaitGroup){

	// 	myChan <- 7
	// 	myChan <- 11
		
	// 	close(myChan)
	// 	wg.Done()
	// }(myChan, wg)



	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup){
		val, isChanOpen := <- myChan

		if isChanOpen {
			fmt.Println("My channel value : ", val)
		}
		wg.Done()
	}(myChan, wg)

	go func(ch <-chan int, wg *sync.WaitGroup){

		myChan <- 0
		close(myChan)
		
		wg.Done()
	}(myChan, wg)



	// wg.Add(2)
	// // Receive Only
	// go func(ch <-chan int, wg *sync.WaitGroup){
	// 	close(myChan) // preventing this error adding <- 
	// 	val, isChanOpen := <- myChan

	// 	if isChanOpen {
	// 		fmt.Println("My channel value : ", val)
	// 	}
	// 	wg.Done()
	// }(myChan, wg)

	// // Send Only
	// go func(ch chan<- int, wg *sync.WaitGroup){

	// 	myChan <- 0
	// 	close(myChan)
		
	// 	wg.Done()
	// }(myChan, wg)

	wg.Wait()
}