package main

import (
	"fmt"
	"reflect"
)

func main() {
	copyFunc()
}

func set_create() {
	var x = []int{1, 2, 3}    // dont set size
	x = append(x, 4, 5, 6, 7) // add element IN THE BACK!
	y := []int{10, 20, 30}
	var xy = append(x, y...) // y... <- unpack slice

	makeSlice := make([]int, 10) //type,length,cup
	fmt.Println(cap(makeSlice), len(makeSlice), reflect.TypeOf(makeSlice))
	fmt.Println(xy)
	fmt.Print(x)
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

func sliceSliceWithGeneralMemory() {
	//USE general peace of memory (if we create slice based on slice)
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]

	x[1] = 20
	y[0] = 10
	z[1] = 30

	fmt.Println("x= ", x)
	fmt.Println("y= ", y)
	fmt.Println("z= ", z)

}
func notPrettyCodeWitAppend() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x: ", x, "y: ", y)
	//x:  [1 2 30 4] y:  [1 2 30]
}
func sliceFromArray() {
	x := [4]int{1, 2, 3, 4}
	y := x[:2]
	x[0] = 10
	fmt.Println("x: ", x, "y: ", y)
}
func copyFunc() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println("x: ", x, "y: ", y, "num: ", num)
}
