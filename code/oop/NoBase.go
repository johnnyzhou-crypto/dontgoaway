package main

import (
	"log"
	"reflect"
)

// START_OMIT
type Base struct{}

type Child struct {
	Base
}

func (b Base) printReceiver() {
	log.Printf("I am: %s", reflect.TypeOf(b).Name())
}

func main() {
	a := Base{}
	b := Child{Base{}}

	a.printReceiver() //I am Base
	b.printReceiver() //I am ???
}

// END_OMIT
