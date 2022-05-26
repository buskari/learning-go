package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	phone string
}

// when a variable is declared but not assigned, Go assign a "zero" value ("" for string, 0 for int and float, false for bool)
func main() {
	p := person{
		firstName: "Andris",
		lastName:  "Buscariolli",
		contact: contact{
			email: "andris.contact@gmail.com",
			phone: "+55(11)99999999",
		},
	}

	p.updateName("Fulano")
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
