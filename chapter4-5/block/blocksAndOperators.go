package main

import (
	"fmt"
	"math/rand"
)

func main() {
	goTo()
}
func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	t := 1
	for t < 5 {
		t++
		fmt.Println(t)
	}

}
func forRange() {
	evenVals := [5]string{"1", "534", "543", "34534", "534534"}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func exampleScope() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}
func goTo() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("success")

done:
	fmt.Println("done")
	fmt.Println(a)

}
