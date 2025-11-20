*append* make x2 cup :(

```go
 x := make([]int, 5) // [0,0,0,0,0]([]int,5,5)
//if we use append(x,1) we get
[0, 0, 0, 0, 0, 1] //([]int,6,10) ( CUP x2)

```

in Slice Slicewe include first value but exclude second value

```go
x := []int{1, 2, 3, 4}
y := x[:2] //[1 2]
z := x[1:] //[2 3 4]
d := x[1:3] //[2 3]
e := x[:] //[1 2 3 4]
```
When we create slice based on SLICE OR ARRAY we create a new var which use general peace of memory
```go
x := []int{1, 2, 3, 4}
y := x[:2]
z := x[1:]

x[1] = 20
y[0] = 10
z[1] = 30

fmt.Println("x= ", x)
fmt.Println("y= ", y)
fmt.Println("z= ", z)
//x=  [10 20 30 4]
//y=  [10 20]
//z=  [20 30 4]

```

Append - make misunderstanding in code (slice based on slice)
```go

x := []int{1, 2, 3, 4}
y := x[:2]
fmt.Println(cap(x), cap(y))
y = append(y, 30)
fmt.Println("x: ", x, "y: ", y)
//x:  [1 2 30 4] y:  [1 2 30]

```
copy(Y,X) copy X to Y
return cont copied elements
```go

x := []int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)
fmt.Println("y: ", y, "num: ", num) 
// y:  [1 2 3 4] num:  4

```
When a slice is copied or passed to a function, Go creates a new slice header (pointer, length, capacity), but the underlying array is NOT copied. Both slices refer to the same underlying memory.