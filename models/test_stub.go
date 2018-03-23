package models

import (
	"fmt"
)

var testInputUsers = []struct {
	name  string
	email string
}{
	{"abe", "abe@example.com"},
	{"bobby", "bobby@example.com"},
	{"carl", "carl@example.com"},
	{"dick", "dick@example.com"},
	{"erick", "erick@example.com"},
	{"fredric", "fredric@example.com"},
	{"gabriel", "gabriel@example.com"},
	{"hana", "hana@example.com"},
	{"ibe", "ibe@example.com"},
	{"jojo", "jojo@example.com"},
}

func createTestUsers() []*User {
	var users []*User
	for _, input := range testInputUsers {
		newUser, err := CreateUser(input.name, input.email)
		if err == nil {
			users = append(users, newUser)
		} else {
			fmt.Println(err)
		}
	}

	return users
}
