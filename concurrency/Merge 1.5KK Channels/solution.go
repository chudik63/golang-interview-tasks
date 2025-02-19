package main

import (
	"runtime"
	"sync"
)

var maxParallel int = runtime.NumCPU() * 2

func main() {
	count := 1500000
	chans := make([]chan int, count)

	for i, ch := range chans {
		ch = make(chan int, 1)
		ch <- i
		close(ch)
		chans[i] = ch
	}

	for item := range mergeCh(chans...) {
		_ = item
	}
}

func mergeCh(chs ...chan int) chan int {
	merged := make(chan int)

	poolCh := make(chan chan int)

	wg := &sync.WaitGroup{}

	go func() {
		defer close(poolCh)
		for _, ch := range chs {
			poolCh <- ch
		}
	}()

	wg.Add(maxParallel)
	for i := 0; i < maxParallel; i++ {
		go func() {
			defer wg.Done()
			for ch := range poolCh {
				for v := range ch {
					merged <- v
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
