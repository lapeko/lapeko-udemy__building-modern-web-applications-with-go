package main

import "fmt"

type user struct {
	firstName string
	age       uint8
}

func (u *user) getName() string {
	return u.firstName
}

func main() {
	user1 := user{
		firstName: "Test",
	}
	fmt.Println(user1)
}
