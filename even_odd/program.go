package main

import "fmt"

func main() {
	var array1 []int

	for i := 0; i < 12; i++ {
		array1 = append(array1, i)
	}

	for i := 0; i < 12; i++ {
		if array1[i]%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
