package main

import (
	"encoding/json"
	"fmt"
)

type MyString string

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u User) GetName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user := User{Name: "Pedro Pessoa", ID: 10}
	fmt.Println(user)
	user.GetName()
	user.SetName("Guilherme Chaves")
	fmt.Println(user)

	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
