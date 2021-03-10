package main

import "fmt"

func main() {
	a := 5
	if a > 10 {
		fmt.Println("a > 10")
	} else if a == 5 {
		fmt.Println("a == 5")
	} else {
		fmt.Println("else")
	}

	b := true
	if x := "Cool"; b {
		fmt.Println(x)
	}
}
