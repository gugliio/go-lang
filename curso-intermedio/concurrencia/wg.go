package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}

func doSomething(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start", index)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
