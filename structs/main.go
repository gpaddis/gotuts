package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 95866,
		},
	}

	// jimPointer := &jim // pointer to jim
	// Person is automatically converted to pointerToPerson
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// alex := person{firstName: "Alex", lastName: "Anderson"}
// fmt.Println(alex)

// var alex person

// alex.firstName = "Alex"
// alex.lastName = "Anderson"

// fmt.Println(alex)
// fmt.Printf("%+v", alex) // Print field names
