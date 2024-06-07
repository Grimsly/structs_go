package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// Can also rely on order of struct field
	// alex := person{"Alex", "Anderson"}

	// When you create variable in Go and do not assign a value to it
	// Go will auto assign a zero value depending on the type (e.g. string will be "")
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)

	alex.print()

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// & is an operator that gives the memory address this variable is pointing at
	// jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print()

}

// Go always passes by value, so it can't modify a receiver/argument passed in if a pointer is not passed in
// Exceptions are reference types which don't point to the value you want modified,
// but it's value contains a pointer to the actual value
// For example, a slice contains a value that points to an array, copying it would just copy the pointer to the array
// so modifying a copy of a slice would still modify the array
// *person is a type description. It's a pointer to a person type.
func (p *person) updateName(newFirstName string) {
	// This * is an operator and is different from the * above
	// This retrieves the value within the pointer
	// (*p).firstName = newFirstName
	// Go syntactic sugar that is the same as the above
	// Since firstName belongs to p (person) and newFirstName is assignable to firstName
	// Go will just assume that the value is to be replaced
	// https://tour.golang.org/moretypes/4
	p.firstName = newFirstName
	// Turn address into value with *address
	// Turn value into address with &value
}

func (p person) print() {
	// %+v will print out all field name and their values for a given struct
	fmt.Printf("%+v\n", p)
}