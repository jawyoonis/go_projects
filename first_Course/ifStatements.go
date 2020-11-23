package main

type Users struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := Users{
		ID:        2,
		FirstName: "Jamal",
		LastName:  "Yonis",
	}

	u2 := Users{
		ID:        1,
		FirstName: "khalid",
		LastName:  "Adnan",
	}

	if u2.ID == u1.ID {
		println("same user ids")
	} else if u1.FirstName == u2.FirstName {
		println("same first names")
	} else {
		println("different user ids and first names")
	}

}
