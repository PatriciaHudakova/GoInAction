package main

type user struct { //define a struct called user
	name       string
	email      string
	ext        int
	privileged bool
}

var peter user //peter is now a type user initialised to zero

var lisa = user{ //or user {}
	name:       "lisa",
	email:      "lisa@email.com",
	ext:        123,
	privileged: true,
}

var pat = user{"peter", "peter@email.com", 345, false} //declaring struct without fields

type admin struct { //declare field based on other struct types
	person user
	level  string
}

var patrick = admin{
	person: user{
		name:       "patrick",
		email:      "patrick@email.com",
		ext:        999,
		privileged: true,
	},
	level: "super",
}

type Duration int64
