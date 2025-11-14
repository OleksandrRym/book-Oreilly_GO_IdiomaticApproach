ok - boolean type which say - map contain element or not
```go
m := map[string]int{
"hello": 5,
"world": 0,
}
v, ok := m["hello"]
fmt.Println(v, ok) 
delete(m, "hello")
v, ok = m["hello"]
fmt.Println(v, ok)
//5 true
//0 false
```
WE HAVENT SET!!!!!!!!!!!!!!!!!!!!!!!!
actualy in java we also use set with object(key) in map, so its can be
```go

```
