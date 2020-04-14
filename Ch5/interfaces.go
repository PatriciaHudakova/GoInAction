package main

import "fmt"

type notifier interface { //defines notification type behavior
	notify()
}

type NotificationUser struct { //user defines a user in the program
	name string
	email string
}

func (u *NotificationUser) notify() { //interface with a pointer receiver
	fmt.Println("Sending user email to ", u.name, u.email)
}

type NotificationAdmin struct {
	name string
	email string
}

func (a *NotificationAdmin) notify() { //interface with a pointer receiver
	fmt.Println("Sending user email to ", a.name, a.email)

func main() {
	bill := NotificationUser{"Bill", "bill@email.com"}  //create value of type user
	sendNotification(&bill)

	lisa := NotificationAdmin{"Lisa", "lisa@email.com"}  //create value of type admin
	sendNotification(&lisa)
}

func sendNotification(n notifier) {
	n.notify()
}