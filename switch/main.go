package main

import "fmt"

func main() {
	a := "Robson"

	switch a {
	case "Bob":
		fmt.Println("Hey Bob")

	case "Maria":
		fmt.Println("Hey Maria")

	default:
		fmt.Println("I did not find your name!")
	}

}
