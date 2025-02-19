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

func main() {
	count := int(1e7)

	codes := &Codes{
		h:  make(map[int]int),
		mu: &sync.Mutex{},
	}

	urls := make([]string, count)

	sema := make(chan struct{}, maxParallel)

	wg := &sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		sema <- struct{}{}

		go func() {
			defer func() {
				<-sema
				wg.Done()
			}()

			code := sendRequest(urls[i])
			codes.inc(code)
		}()
	}

	wg.Wait()

	close(sema)

	fmt.Println(codes.h)
}

func sendRequest(url string) (code int) {
	codes := []int{200, 404, 500, 201, 505, 401, 403}
	return codes[rand.Intn(len(codes))]
}
