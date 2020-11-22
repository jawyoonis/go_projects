package main

import "fmt"

func main() {

	type user struct {
		ID        int
		firstName string
		lastName  string
	}

	var u user
	u.ID = 1
	u.firstName = "jamal"
	u.lastName = "Aw Yonis"
	fmt.Println(u.ID)
	fmt.Println(u.firstName, u.lastName)
}
