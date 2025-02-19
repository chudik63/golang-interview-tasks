```go
package main

import "fmt"

type C chan C

func main() {
	var c = make(C, 1)
	c <- c
	for i := 0; i < 1000; i++ {
		select {
		case <-c:
		case <-c:
			c <- c
		default:
			fmt.Println(i)
			return
		}
	}
}
```
# Select-Case
What is happening here?