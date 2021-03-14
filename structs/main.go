package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	// method 1
	name := person{firstname: "Mahan", lastname: "R"}
	fmt.Println(name)

	// method 2
	var name1 person
	fmt.Println(name1)
	fmt.Printf("%+v", name1)

	name1.firstname = "Under"
	name1.lastname = "taker"
	fmt.Printf("%+v", name1)
}
