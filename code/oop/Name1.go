package main

import "fmt"

// START_OMIT
type Name string

func (n Name) Greeting() string {
	return fmt.Sprintf("Hello %s", n)
}

func main() {
	var n Name = "Thomas"
	fmt.Println(n.Greeting())
}

// END_OMIT
