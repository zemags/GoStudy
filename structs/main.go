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

	jim := person {
		firstName: "Jim",
		lastName: "K",
		contact: contactInfo{
			email: "jim@jim.com",
			zipCode: 123456,
		},
	}

	fmt.Printf("%+v", jim)
}
