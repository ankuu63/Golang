package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("hey channels ")

	wg := &sync.WaitGroup{}
	wg.Add(2)

	mych := make(chan int, 4)

	go func(wg *sync.WaitGroup, ch chan int) {

		fmt.Println(<-mych)

		//fmt.Println(<-mych)
		wg.Done()
	}(wg, mych)

	go func(wg *sync.WaitGroup, ch chan int) {
		mych <- 5
		mych <- 6

		wg.Done()
	}(wg, mych)

	wg.Wait()

	fmt.Println("channels has succesfully created ")
	fmt.Println("channels  ")
}
