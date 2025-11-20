package main

import (
	"fmt"
	"time"
)

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

}

type Person struct {
	age      int
	name     string
	lastname string
}
type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}
func (c Counter) String() string {
	return fmt.Sprintf("total: %s, last update  %s ", c.total, c.lastUpdate)
}

// method receiver
func (p Person) seyHello() string {
	return fmt.Sprintf("%s %s - %s", p.lastname, p.name, p.age)
}
