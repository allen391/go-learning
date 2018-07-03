package main

import (
	"fmt"
)

func arr1() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 5, 6, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arr1()
	arr4 := [...]int{2, 3, 4, 2, 1}
	for _, v := range arr4 {
		fmt.Println(v)
	}
	printArray(arr4)
	fmt.Println(arr4)
}
