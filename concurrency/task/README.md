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
# Task 1
Есть фугкия, вызов которой стоит денег - `requestThatCostsMoney`. Ее внутренности не важны. Главное то, что каждый запрос к ней стоит нам денег. Запускаются множество горутин, каждая из которых обращается к этой функции. Задача - оптимизировать вызов так, чтобы из нескольких горутин с одинаковыми `userID` в запросе, которые вызывают функцию параллельно, работала только одна и делилась полученным ответов с другими. Таким образом мы не будет терять деньги на лишние вызовы функции с одним и тем же запросом одновременно.

В коде уже заготовлена структура с методом, который является оберткой над функцией `requestThatCostsMoney`, в  нем должна содержаться вся логика.