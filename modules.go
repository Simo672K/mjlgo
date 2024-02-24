package main

import (
	"fmt"
	"learning-go/types"
)

func modules() {
	newUser := types.User{
		Username: "this is a username",
		Age:      24,
	}

	fmt.Println("user", newUser)
}
