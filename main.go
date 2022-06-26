package main

import "fmt"


type User struct {
	FullName string
	Gender   string
	Age      int
	Email    string
}

func main() {
	var users = []User{
		{
			FullName: "Azman",
			Gender:   L,
			Age:      23,
			Email:    "azman@mail.com",
		},
		{
			FullName: "Dian",
			Gender:   P,
			Age:      20,
			Email:    "dian@mail.com",
		},
	}

	fmt.Println(users)
}
