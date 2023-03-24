package main

import (
	"fmt"
	"sync"
)

//consume, provider

func provider(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	//close(ch)
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		fmt.Printf("%v\n", c)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go provider(ch, &wg)
	go consume(ch, &wg)
	wg.Wait()
}
