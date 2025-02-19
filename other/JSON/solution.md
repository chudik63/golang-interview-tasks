# Solution

1) You need to know that the `%#v` specifier outputs the variable in a syntactically correct format, i.e. in a form in which it can be written in the code. Output: `main.Data{X:1, y:"two"}`
2) You need to know that private variables are not serialized in `json`. Output: `{"x":1}`
3) When `Unmarshal` is used, the `out` variable will not have a value for the `y` field, since it was not in `json`, so it will take a zero value - an empty string. Output:
`main.Data{X:1, y:""}`
