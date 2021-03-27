package main

import "fmt"

const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(s string, language string) string {
	if s == "" {
		s = "Mahan"
	}

	return greetingPrefix(language) + s
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Mahan", ""))
}

// How to run - go run main.go
// Package main - {package == project == workspace}
// * Executable - Generates a executable file that can be run
// * Reusable - code used as "helpers", reusable code.
//
