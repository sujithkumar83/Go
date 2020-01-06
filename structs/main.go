package main

import "fmt"

// Embed on struct within another

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// 2 ways of declaring a struct
// printf demo: %v substitutes the variable inside the string to be printed

func main() {

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "aa@abc.com",
			zipCode: 94000,
		},
	}
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	//alexPointer := &alex
	//alexPointer.updateName("jimmy")

	alex.updateName("jimmy")

	alex.print()

}

// Reciever funciton with Struct
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newfirstName string) {
	(*pointerToPerson).firstName = newfirstName
}
