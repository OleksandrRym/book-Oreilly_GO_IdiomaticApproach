package main

import "fmt"

func main() {
	createInstance()
}
func createAnonymStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.pet = "pet boba"
	person.age = 12
} // WHAAAAAT

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

// typ
