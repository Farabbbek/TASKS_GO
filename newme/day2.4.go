package main

import "fmt"

func Slicewithslice() {
	s := []int{45, 66, 77, 88, 99}
	//fmt.Println(s)
	s2 := s[0:2]
	s2 = append(s2, 500)
	fmt.Println(s)
	fmt.Println(s2)

}
