package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	current := "0"

	printFunc := func(name, id, next string) {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			mu.Lock()
			for current != id {
				cond.Wait()
			}

			fmt.Print(name)
			if i == 10-1 {

				current = next
				cond.Broadcast()
			}
			mu.Unlock()
		}
	}

	wg.Add(3)

	go printFunc("A", "0", "1")
	go printFunc("B", "1", "2")
	go printFunc("C", "2", "0")

	wg.Wait()

	fmt.Println("\n done")

}
