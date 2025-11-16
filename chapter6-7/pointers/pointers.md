pointers - its just var which contains address \
&val - operator get address\
*val -dereference operator

It’s better to create a function that returns a fully-initialized instance, rather than receiving a pointer and mutating it.

```go
func makeFooGoodCase() (Foo, error) {
f := Foo{
"1",
1,
}
return f, nil
}


func makeFooBadCase(f *Foo) error {
f.f1 = "1"
f.f2 = 1
return nil
}

```
