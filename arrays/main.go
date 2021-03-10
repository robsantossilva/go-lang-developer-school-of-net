package main

import "fmt"

func main() {
	var x [10]int //10 zero positions
	fmt.Println(x)
	fmt.Println(len(x))

	for i, _ := range x {
		x[i] = i * 2
	}

	fmt.Println(x)

	//#############################################
	fmt.Println("//#############################################")

	var y [10]string
	fmt.Println(y)
	fmt.Println(len(y))

	for i, _ := range y {
		y[i] = "A"
	}

	fmt.Println(y)

	//#############################################
	fmt.Println("//#############################################")

	z := [5]int{4, 56, 1, 4, 78}
	fmt.Println(z)
}
