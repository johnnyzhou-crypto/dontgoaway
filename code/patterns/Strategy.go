package main

import "fmt"

// START_DEF_OMIT

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

// END_DEF_OMIT

// START_IMPL_OMIT

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

// END_IMPL_OMIT

// START_RUN_OMIT
func main() {
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5))

	multi := Operation{Multiplication{}}
	fmt.Println(multi.Operate(3, 5))
}

// END_RUN_OMIT
