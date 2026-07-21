package main

import (
	"fmt"
)

func ArraySlice() {
	//array
	arr := [5]int{73, 98, 86, 61, 96}

	fmt.Println(arr)
	fmt.Println(arr[3])

	arr[4] = 100
	fmt.Println(arr)

	//slice
	s := []int{73, 98, 86, 61, 96}
	fmt.Println("Длина =", len(s), "вместимость =", cap(s))
	fmt.Println(s)
	s[2] = 33
	s = append(s, 999)
	fmt.Println(s)
	fmt.Println("Длина =", len(s), "вместимость =", cap(s))

	//slice с фиксированной вместимостью
	s2 := make([]int, 5, 500)
	fmt.Println(s2)
	fmt.Println("Длина =", len(s2), "вместимость =", cap(s2))

	//срезы
	//s := []int{73, 98, 86, 61, 96}
	s3 := s[0:2:2]
	fmt.Println(s3)
	fmt.Println("Длина =", len(s3), "вместимость =", cap(s3))

}
