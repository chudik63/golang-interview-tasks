package main

import "runtime"

var maxParallel int = runtime.NumCPU() * 2

func main() {
	// logic

	for item := range mergeCh(ch1, ch2, ch3...) {
		// reading from merged channel
	}
}

func mergeCh(chs ...chan string) stirng {
	poolCh := make(chan chan string, maxParallel)

	go func() {
		for _, ch := range chs {
			poolCh <- ch
		}
	}
}
