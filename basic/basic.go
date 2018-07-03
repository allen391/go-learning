package main

import (
	"fmt"
	"math"
)

var (
	aa = 3
	ss = "string"
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)

}

func variableShorter() {
	//only allow to declare inside the block scope
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, s, c)
}

func constName() {
	const (
		filename = "abc.text"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb)
}

func main() {
	fmt.Println("hello world")
	variable()
	variableInitialValue()
	constName()
	enums()
}

//datatypes
//bool, string
//int
//byte, rune
//float32, float64, complex64, complex128
