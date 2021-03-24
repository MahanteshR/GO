package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i)
		//time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fun("Direct call")

	go fun("routine call")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous func call")

	time.Sleep(time.Second)
	fmt.Println("done")
}
