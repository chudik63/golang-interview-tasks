# Task 1

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Data struct {
    X int `json:"x"`
    y string `json:"y"`
}

func main() {
    in := Data{1, "two"}
    fmt.Printf("%#v\n", in)

    encoded, _ := json.Marshal(in)
    fmt.Println(string(encoded))

    var out Data
    json.Unmarshal(encoded, &out)
    fmt.Printf("%#v\n", out)
}
```
Что будет выведено в консоль?