package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning about Race Conditions using Go Programming")

	var score = []int{0}

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("R-1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("R-2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("R-3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("R-4")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
}
