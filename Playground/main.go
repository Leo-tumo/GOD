package main

import (
	"fmt"
	"sync"
	"time"
)

const workers = 3

func job(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for w := 1; w <= workers; w++ {

		go job(&wg)
	}

	fmt.Println("Hello World from Go")

	wg.Wait()
	time.Sleep(1 * time.Second)
}
