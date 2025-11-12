package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = []int{1, 2, 3}    // dont set size
	x = append(x, 4, 5, 6, 7) // add element IN THE BACK!
	y := []int{10, 20, 30}
	var xy = append(x, y...) // y... <- unpack slice

	makeSlice := make([]int, 10) //type,length,cup
	fmt.Println(cap(makeSlice), len(makeSlice), reflect.TypeOf(makeSlice))
	fmt.Println(xy)
	fmt.Print(x)
	sliceSlice()
}
func sliceSlice() {
	//include a but exclude b[a:b]
	//USE general peace of memory (if we create slice based on slice)
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x= ", x)
	fmt.Println("y= ", y)

	fmt.Println("z= ", z)
	fmt.Println("d= ", d)
	fmt.Println("e= ", e)

}
