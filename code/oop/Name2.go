package main

import "fmt"

// START_OMIT
type Name string

func (n Name) Greeting() string {
	return fmt.Sprintf("Hello %s", n)
}

type Person struct {
	Name
	Age int
}

func main() {
	p := &Person{Name("David"), 30}
	fmt.Println(p.Greeting())
}

// END_OMIT
