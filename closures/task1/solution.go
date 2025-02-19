package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
