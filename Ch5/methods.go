package main

import "fmt"

type user2 struct {
	name  string
	email string
}

func (u user2) notify() { //works with a copy of a value, value receiver
	fmt.Println("Sending email to ", u.name, u.email)
}

func (u *user2) changeEmail(email string) { //works with an actual value, pointer receiver
	u.email = email
}

func main() {
	Kajal := user2{
		name:  "Kajal",
		email: "kajal@email.com",
	}

	Pat := &user2{
		name:  "Pat",
		email: "pat@email.com",
	}
	Kajal.notify() //values of type can be used to call methods
	Pat.notify()   //pointer of type can be used to call methods

	Kajal.changeEmail("kajal@newemail.com")
	Kajal.notify()

	Pat.changeEmail("pat@newemail.com")
	Pat.notify()
}
