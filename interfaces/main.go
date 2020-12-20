package main

import "fmt"

type bot interface { // interface type
	getGreeting() string // if any func in program get getGreeting that
	//returns a string - now promoted to be type 'bot' (they are included on
	//type 'bot')
}

type englishBot struct{} // concrete type
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // a bite like a python def
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // englishBot as a receiver
	// some logic
	return "Hi"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
