package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}

func main() {
	var s []int // zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	printSlice(s)

	s1 := []int{2, 4, 6, 7}
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("delete element in slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)

}
