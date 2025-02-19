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
# Task 1
В цикле на каждой итерации запускается горутина, которая захватывает переменную `i`. Что будет выведено в консоль?