package main

import (
	"fmt"
)

type User struct {
	FullName string
	Gender   string
	Age      int
}

func main() {
	var users = []User{
		{
			FullName: "Azman",
			Gender:   L,
			Age:      23,
		},
		{
			FullName: "Dian",
			Gender:   P,
			Age:      20,
		},
	}

	fmt.Println(users)
}
