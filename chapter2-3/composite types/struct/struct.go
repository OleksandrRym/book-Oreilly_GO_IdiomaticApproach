package main

import "fmt"

func main() {

	p1 := fp{
		11,
		"sadcxc",
	}
	p2 := struct {
		age  int8
		name string
	}{11,
		"sadcxc"}
	fmt.Println(p1 == p2)
}
func createAnonymStruct() {
	//Define and setup
	pet := struct {
		name string
		age  int64
	}{
		"Oleg",
		22,
	}
	// Good for murshaling/unmarshaling
	fmt.Println(pet)
}

func createInstance() {
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
	oleg.age = 23 // can set or get
	fmt.Println(julia)
	fmt.Println(oleg)
}

type person struct {
	name string
	age  int
	pet  string
}
type fp struct {
	age  int8
	name string
}
