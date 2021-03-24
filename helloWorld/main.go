package main

import "fmt"

const prefix = "Hello, "

func Hello(s string) string {
	if s == "" {
		s = "Mahan"
	}
	return prefix + s
}

func main() {
	fmt.Println(Hello("Mahan"))
}

// How to run - go run main.go
// Package main - {package == project == workspace}
// * Executable - Generates a executable file that can be run
// * Reusable - code used as "helpers", reusable code.
//
