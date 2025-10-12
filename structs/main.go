package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	age       int
	gender    string
	contact   contactInfo
}

func main() {
	pranav := person{
		firstName: "pranav",
		lastName:  "V",
		age:       22,
		gender:    "Male",
		contact: contactInfo{
			email: "jainvpranav@gmail.com",
			phone: "8618994561",
		},
	}
	pranav.display()
	pranav.updateFirstName("Pranav")
	pranav.display()
}

func (p person) display() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateFirstName(nv string) {
	(*pointerToPerson).firstName = nv
}
