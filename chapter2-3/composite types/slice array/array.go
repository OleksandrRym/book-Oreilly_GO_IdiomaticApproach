package main

import "fmt"

func main() {
	var x = [3]int{1, 2, 3}
	var y = [3]int{1, 2: 3, 1: 2} // u can set value by (index: value)
	var initByLiteral = [...]int{1, 2, 3}
	fmt.Print(initByLiteral == x, initByLiteral == y, x == y) // check value

	fmt.Println(len(initByLiteral))
}
