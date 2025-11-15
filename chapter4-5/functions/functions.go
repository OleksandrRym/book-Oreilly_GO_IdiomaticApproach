package main

import "fmt"

func main() {
	showAnon()
}
func showAnon() {
	for i := 1; i <= 5; i++ {
		func(j int) {
			fmt.Println(j)
		}(i)
	}
}

type opFuncType func(int, int) int
