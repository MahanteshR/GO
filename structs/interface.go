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

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (englishBot) getGreeting() string {
	return "Hi"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
