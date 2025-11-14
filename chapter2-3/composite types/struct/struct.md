struct - composite data type that groups together multiple fields under a single name.
```go
type person struct {
	name string
	age  int
	pet  string
}
```
struct u can init with name or fields or not
```go

julia := person{
    "julia",
    11,
    "tyzik",
}

oleg := person{
	name: "oleg",
    age:  22,
    pet:  "howrah",
}
```
Anonymous struct — a struct defined without a type name, usually declared directly inside a variable with its fields specified inline.
```go

pet := struct {
    name string
    age  int64
}{
    "Oleg",
    22,
} // Good for murshaling/unmarshaling
```
Marshaling - process convert struct go -> json,xml etc\
Demarshaling/Unmarshaling - process convert json -> struct go

EQUALS
To compare two structs, the structs must have the same type, the same fields, and the same field order.
```go
type person struct {
    age  int8
    name string
}
p1 := person{
    11,
    "sadcxc",
}
p2 := struct {
    age  int8
    name string
}{
	11,
    "sadcxc"}
fmt.Println(p1 == p2)// true

```
but we can use anon struct for fast compare