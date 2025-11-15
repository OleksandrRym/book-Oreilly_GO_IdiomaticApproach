variable lives only in block {} code 
```go
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	//10
	//5
    //10
```
universe block -  literally the biggest scope in go
For Range - returned 2 var ( index and values)
copied values ( u don't can change element in array - inside loop)
```go
	evenVals := [5]string{"1", "534", "543", "34534", "534534"}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
```
