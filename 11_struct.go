package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(user []user, name string) (value *user, err error) {
	for _, u := range user {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	users := []user{{"shanliao", "0420"}, {"king", "333"}}
	u, err := findUser(users, "king")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name, u.password)

	if u, err := findUser(users, "none"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*u)
	}
}
