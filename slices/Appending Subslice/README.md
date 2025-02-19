```go
func main() {
	slice := make([]int, 3, 4)
	appendingSlice(slice[:1])
	fmt.Println(slice)
}

func appendingSlice(slice []int) {
	slice = append(slice, 1)
}
```
# Appending Subslice
What will be printed to the console?