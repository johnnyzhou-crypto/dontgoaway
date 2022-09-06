package main

import "fmt"

// START_OMIT
type Person struct {
	Name string
	Age  int
}

func (p *Person) Greeting() {
	fmt.Println("Hello,", p.Name, ". you're", p.Age, "years old")
}
func (p *Person) grow(years int) {
	p.Age += years
}

func main() {
	p := Person{"Thomas", 10}
	p.Greeting()
	p.grow(1)
	p.Greeting()
}

// END_OMIT
