```go
var reqs int

type MoneyOptimization struct {
    // your code here
}

func main() {
	mo := MoneyOptimization{}
	userID := "vasya"

	wg := &sync.WaitGroup{}

	c := 1000

	wg.Add(c)

	for i := 0; i < c; i++ {
		go func() {
			res := mo.RequestThatCostsMoney(userID)
			if res != "some" {
				fmt.Println("!= some")
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(reqs)
}

func requestThatCostsMoney(userID string) (response string) {
	reqs++
	fmt.Println("request ", userID)
	return "some"
}

func (o *MoneyOptimization) RequestThatCostsMoney(userID string) (response string) {
    // your code here
}

```
# Money Optimization
There is a function that costs money to call - `requestThatCostsMoney`. Its insides are not important. The main thing is that each request to it costs us money. Many goroutines are launched, each of which calls this function. The task is to optimize the call so that out of several goroutines with the same `userID` in the request that call the function in parallel, only one works and shares the received response with the others. This way, we will not lose money on unnecessary function calls with the same request at the same time.

The code already has a structure with a method that is a wrapper over the `requestThatCostsMoney` function, it should contain all the logic.