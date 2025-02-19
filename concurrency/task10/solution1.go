package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

var maxParallel int = runtime.NumCPU() * 2

type Codes struct {
	h  map[int]int
	mu *sync.Mutex
}

func (c *Codes) inc(code int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.h[code]++
}

func worker(reqs chan string, codes *Codes, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range reqs {
		code := sendRequest(url)
		codes.inc(code)
	}
}

func main() {
	count := int(1e7)

	codes := &Codes{
		h:  make(map[int]int),
		mu: &sync.Mutex{},
	}

	urls := make([]string, count)

	requests := make(chan string, count)

	wg := &sync.WaitGroup{}

	go func() {
		for i := 0; i < maxParallel; i++ {
			wg.Add(1)
			go worker(requests, codes, wg)
		}
	}()

	for i := 0; i < count; i++ {
		requests <- urls[i]
	}
	close(requests)

	wg.Wait()

	fmt.Println(codes.h)
}

func sendRequest(url string) (code int) {
	codes := []int{200, 404, 500, 201, 505, 401, 403}
	return codes[rand.Intn(len(codes))]
}
