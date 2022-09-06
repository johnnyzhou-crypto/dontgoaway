package main

import "fmt"

// START_OMIT
type Person struct {
	Name string
	Age  int
}

func (p *Person) greeting() string {
	return fmt.Sprintf("Hello, %s", p.Name)
}

func main() {
	var p Person
	fmt.Println(p.greeting())
}

// END_OMIT
