package main

import "fmt"

func xpto(a *int) int {
	*a = *a + 1
	return *a
}

func main() {

	b := 10
	fmt.Println(b)
	fmt.Println(xpto(&b))
	fmt.Println(b)
	fmt.Println("###################")

	x := 10

	y := &x

	*y = 30

	fmt.Printf("y = %T - %v \n", y, y)

	fmt.Printf("x = %T - %v \n", x, x)

	var z *int = &x

	*z = 40

	fmt.Printf("z = %T - %v \n", z, z)

	fmt.Printf("x = %T - %v \n", x, x)
}
