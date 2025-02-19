package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	X int    `json:"x"`
	y string `json:"y"`
}

func main() {
	in := Data{1, "two"}
	fmt.Printf("%#v\n", in) // main.Data{X:1, y:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // {"x":1}

	var out Data
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // main.Data{X:1, y:""}
}
