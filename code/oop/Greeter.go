package main

import "fmt"

type Name string

func (n Name) Greeting() string {
	return fmt.Sprintf("Hello %s", n)
}

type Person struct {
	Name
	Age int
}

// START_OMIT
type Greeter interface {
	Greeting() string
}

func (g Greeter) Greeting() string { //Wrong
	return fmt.Sprintf("Hello %s", g.Greeting())
}

// END_OMIT

func sayHello(g Greeter) {
	fmt.Println(g.Greeting())
}

func main() {
	n := Name("Thomas")
	p := Person{"David", 30}
	greeters := []Greeter{n, p}
	for _, g := range greeters {
		sayHello(g)
	}
}
