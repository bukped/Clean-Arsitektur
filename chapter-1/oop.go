package main

import "fmt"

type user struct {
	name string
}

func (u user) getName() string {
	return u.name
}
func newUser(name string) user {
	return user{name: name}
}
func main() {
	user := newUser("Zaky Muhammad Yusuf")
	fmt.Println(user.getName()) // John Lennon
}

// https://medium.com/@rezaindra/mengenal-konsep-object-oriented-pada-golang-b28e6e2b183c
