package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct { // new custom type of person is oing to be a struct
	firstName string
	lastName string
	contact contactInfo //custome type
}

func main()  {
	// john := person{firstName: "John", lastName: "Doe"} // fristName and lastName
	// john := person{"John", "Doe"} // fristName and lastName, in this var order is important
	// var john person // initiate struct with 'zero' value, for firstName and LastName will be empty string ""

	// john.firstName = "John"
	// john.lastName = "Doe"

	jim := person{
		firstName: "Jim",
		lastName: "K",
		contact: contactInfo{
			email: "jim@jim.com",
			zipCode: 123456,
		},
	}

	// jim := person {	// second var to declare if type person 'contact contactInfo' replace to 'contactInfo'
	// 	firstName: "Jim",
	// 	lastName: "K",
	// 	contactInfo: contactInfo{
	// 		email: "jim@jim.com",
	// 		zipCode: 123456,
	// 	},
	// }

	jimPointer := &jim  //Give memory(RAM) address of the value jim, reference to memery address

	jimPointer.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) { // *person means working with a pointer to a person
	(*pointerToPerson).firstName = newFirstName  // (*pointerToPerson) means take this memory address and give a value of this address, our case its person{firstName: "Jim", ..}, manipulate the value the pointer is referencing
}

func (p person) print() { // recivier
	fmt.Printf("%+v", p)
}