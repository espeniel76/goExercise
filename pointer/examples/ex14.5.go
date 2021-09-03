package examples

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	// var u = User{name, age}
	u := &User{name, age}
	return u
}

func Ex14_5() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
