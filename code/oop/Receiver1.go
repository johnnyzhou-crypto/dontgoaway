package main

import "fmt"

// START_OMIT
type Person struct {
	Name string
	Age  int
}

func greeting(p Person) string {
	return fmt.Sprintf("Hello, %s", p.Name)
}

func (p Person) greeting() string {
	return fmt.Sprintf("Hello, %s", p.Name)
}

func main() {
	p := Person{"David", 30}
	fmt.Println(greeting(p))
	fmt.Println(p.greeting())
}

// END_OMIT
