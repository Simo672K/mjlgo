package main

import "fmt"

var (
	myVar        = "this is a string variable"
	myNumVar int = 10
)

const thisIsMyConstent = 10 // this is how constants are declared in golang

func Vars() {
	// or we can declar and infer vals
	myOtherNumVar := 15
	fmt.Printf("Printing 'myOtherNumVar': %d\n", myOtherNumVar)
}
