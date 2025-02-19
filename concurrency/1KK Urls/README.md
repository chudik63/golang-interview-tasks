# Task 10
```go
package main

func main() {
	urls := make([]string, 1000000000)
	codes := make(map[int]int)

	_ = urls
	_ = codes
}

```
# 1KK Urls
There are a billion links. Concurrency must be used to efficiently store in a hash table the number of links that returned a certain code:
```
    200 - 3
    404 - 1
    500 - 100
```
