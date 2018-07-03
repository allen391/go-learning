package main

import (
	"fmt"
	"io/ioutil"
)

//switch
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score < 100:
		g = "a"
	default:
		panic(fmt.Sprintf("wrong score: %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
	fmt.Println(
		grade(0),
		grade(59),
		grade(99),
	)
}
