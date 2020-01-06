package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}
func (englishBot) getGreeting() string {
	// Very custom logic here
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic here
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
