package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 4, 5, 6}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[0:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
	//extent slice
	fmt.Println("extending slice")
	s1 := arr[2:6]
	s2 := arr[3:5]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
}

// slice is a view of the array
