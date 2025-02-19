```go
package main


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

func mergeCh(chs ...chan string) chan string {
    // channel merging
}
```
# Merge 1.5KK Channels
Merge a large number of channels into one.