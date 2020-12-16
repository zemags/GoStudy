package main

import "fmt"

type person struct { // new custom type of person is oing to be a struct
	firstName string
	lastName string
}

func main()  {
	// john := person{firstName: "John", lastName: "Doe"} // fristName and lastName
	// john := person{"John", "Doe"} // fristName and lastName, in this var order is important
	var john person // initiate struct with 'zero' value, for firstName and LastName will be empty string ""

	john.firstName = "John"
	john.lastName = "Doe"

	fmt.Println(john)
	fmt.Printf("%+v", john)  // print with struct fields
}
