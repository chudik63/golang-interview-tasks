```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		} ()
	}
}

```
# Goroutine
The loop starts a goroutine capturing `i` on each iteraton. What is the output?