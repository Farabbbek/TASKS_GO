package main

import "fmt"

func OddEven() {
	var a int
	fmt.Scan(&a)
	if a%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
