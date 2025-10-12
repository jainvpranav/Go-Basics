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
		firstName: "Pranav",
		lastName:  "V",
		age:       22,
		gender:    "Male",
		contact: contactInfo{
			email: "jainvpranav@gmail.com",
			phone: "8618994561",
		},
	}

	pranav.lastName = "Jain"

	fmt.Println(pranav)
	fmt.Printf("%+v", pranav)
}
