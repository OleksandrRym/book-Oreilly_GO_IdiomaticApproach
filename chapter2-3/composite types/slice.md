*append* make x2 cup :(
```go
 x := make([]int, 5) // [0,0,0,0,0]([]int,5,5)
 //if we use append(x,1) we get
[0,0,0,0,0,1]//([]int,6,10) ( CUP x2)

```
in Slice Slicewe include first value but exclude second value
```go
x := []int{1, 2, 3, 4}
y := x[:2] //[1 2]
z := x[1:] //[2 3 4]
d := x[1:3]//[2 3]
e := x[:]//[1 2 3 4]
```
