package main

import "fmt"

func calculator(a, b int) {

	fmt.Scan(&a, &b)

	fmt.Println(a + b)

}

func main() {
	var a, b int
	calculator(a, b)
}
