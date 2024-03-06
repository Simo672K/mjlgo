package main

import (
	"fmt"
	"time"
)

type Server struct {
	name string
}


func SomeHeavyFunction() string{
	time.Sleep(time.Second * 5)
	return "some heavy stuffs going on here"
}

func MainRoutines() {
	fmt.Println("start of the programme...") 
	go SomeHeavyFunction() // this is what so called go routine, function prefixed with go keyword that executes in the golang schedular (async)
}

