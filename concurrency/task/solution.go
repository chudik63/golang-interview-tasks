package main

import (
	"fmt"
	"sync"
)

type MoneyOptimization struct {
	cache map[string]*G
	mu    *sync.Mutex
}

type G struct {
	ch       chan struct{}
	response string
}

func main() {
	mo := MoneyOptimization{
		cache: make(map[string]*G),
		mu:    &sync.Mutex{},
	}
	userID := "vasya"

	wg := &sync.WaitGroup{}

	c := 1000

	wg.Add(c)

	for i := 0; i < c; i++ {
		go func() {
			res := mo.RequestThatCostsMoney(userID)
			if res != "some" {
				fmt.Println("expected: some, got: ", res)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("total: ", reqs)
}

var reqs int

func requestThatCostsMoney(userID string) (response string) {
	reqs++
	fmt.Println("request: ", userID)
	return "some"
}

func (o *MoneyOptimization) RequestThatCostsMoney(userID string) (response string) {
	o.mu.Lock()
	if cache, ok := o.cache[userID]; ok {
		o.mu.Unlock()

		<-cache.ch
		return cache.response
	}

	o.cache[userID] = &G{ch: make(chan struct{})}
	o.mu.Unlock()

	response = requestThatCostsMoney(userID)

	o.mu.Lock()
	o.cache[userID].response = response
	o.mu.Unlock()

	close(o.cache[userID].ch)

	o.mu.Lock()
	delete(o.cache, userID)
	o.mu.Unlock()

	return response
}
