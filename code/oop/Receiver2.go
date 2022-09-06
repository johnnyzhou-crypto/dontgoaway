package main

import "fmt"

// START_OMIT
type Person struct {
	Name string
	Age  int
}

func (p Person) changeAge(age int) {
	p.Age = age
}

func changeAge(p Person, age int) {
	p.Age = age
}

func main() {
	p := Person{"David", 30}
	p.changeAge(33)
	fmt.Println(p.Age)
	changeAge(p, 35)
	fmt.Println(p.Age)
}

// END_OMIT
