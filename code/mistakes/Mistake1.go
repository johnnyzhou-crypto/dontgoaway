package main

import "fmt"

// START_OMIT
func greatThanZero(arg int) interface{} {
	var result *struct{} = nil
	if arg > 0 {
		result = &struct{}{} //good parameters
	}
	return result //bad parameters
}

func main() {
	if res := greatThanZero(-1); res != nil {
		fmt.Println("Good parameters:", res)
	} else {
		fmt.Println("Bad parameters:", res)
	}
}

// END_OMIT
