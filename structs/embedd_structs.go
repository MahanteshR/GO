package main

import "fmt"

type contactInfo struct {
	address string
	zip     int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo // using only contactInfo is also fine
}

func main() {
	jim := person{
		firstname: "Alex",
		lastname:  "hales",
		contact: contactInfo{ // if only contactInfo is used above, contactInfo: contactInfo{}
			address: "Baker street", zip: 12345,
		},
	}
	// the ',' is necessary after every struct member
	fmt.Printf("%+v", jim)
}
