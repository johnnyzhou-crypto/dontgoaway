package main

import "fmt"

type Country interface {
	Slogan() string
	//History() int
}

type England struct {
	Name string
}

type Germany struct {
	Name string
}

func (eng England) Slogan() string {
	return eng.Name + " Football"
}
func (ger Germany) Slogan() string {
	return ger.Name + " Beer"
}

func SaySomething(country Country) {
	fmt.Println(country.Slogan())
}

func main() {
	var eng Country = England{"England"}
	var ger Country = Germany{"Germany"}

	SaySomething(eng)
	SaySomething(ger)

}
