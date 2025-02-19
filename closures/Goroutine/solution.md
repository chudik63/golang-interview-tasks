# Solution
## < 1.22

First, higly likely nothing. In order to fix this a `WaitGroup` needs to be added, because the loop and `main` will finish faster than the first goroutine starts.
```go
package main

import "fmt"

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			fmt.Println(i)
			wg.Done()
		} ()
	}

	wg.Wait()
}
```

Second, depending on the order in which the goroutines start executing, the output may differ, but the fact is that we will not see all the numbers in order. The answer may be either

```
5
5
5
5
5
```
or, for example,
```
5
5
5
3
3
```
Some goroutines that was called on the `i = 3` iteration managed to start earlier and capture the variable in the middle of the loop.

The problem is that in versions before 1.22 such calls inside a loop captures a variable by pointer. Accordingly, goroutines that started after the loop referenced `i`, which at the end is equal to `5`. 

How to solve it?

1) Pass the variable as a functon argument
```go
package main

import "fmt"

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i)
			wg.Done()
		} (i)
	}

	wg.Wait()
}
```

2) Shadowing
```go
package main

import "fmt"

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			fmt.Println(i)
			wg.Done()
		} ()
	}

	wg.Wait()
}
```

## > 1.22
After adding `WaitGroup` we will see all the digits in the specified range. However, the order of their output cannot be guaranteed, since everything depends on the scheduler:

```
4
1
0
3
2
```