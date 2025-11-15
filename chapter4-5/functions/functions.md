Functional types - It represents any function that takes two int arguments and returns an int.
```go

type opFuncType func (int, int) int
// alias for group 
func calc(op opFuncType, x, y int) int {
    return op(x, y)
}

```
Anonym func -  a function without a name, often used as a value, callback, or closure.

```go
    for i := 1; i <= 5; i++ {
    func (j int) {
         fmt.Println(j)
    }(i)
}


func(a, b int) int {
    return a * b
}(3, 4) // 12
```
