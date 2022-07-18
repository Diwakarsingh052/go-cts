package main

import "fmt"

type user struct {
	name, email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type admin struct {
	user  // embedding  // anonymous field because we are not giving any name to it // anonymous field name would be assigned according to the type name
	roles []string
}

func (a admin) show() {
	fmt.Println(a)
}

type dev struct {
	user
	project string
}

func main() {
	// admin is embedding user struct ,so it can use all its fields and its methods
	a := admin{
		user: user{
			name:  "raj",
			email: "email.com",
		},
		roles: []string{"admin"},
	}
	a.show()
	a.updateEmail("real@email.com")
	a.show()

	a1 := admin{
		user: user{
			name:  "a1",
			email: "a1",
		},
		roles: nil,
	}
	a1.show()

	u := user{
		name:  "user name",
		email: "user email",
	}
	fmt.Println(u)
	a.show()

}
