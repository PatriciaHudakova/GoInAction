package unexported

type user struct { //unexported
	Name  string //exported
	Email string //exported
}

type Admin struct { //exported
	user
	Rights int
}
