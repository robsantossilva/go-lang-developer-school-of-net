package main

import "fmt"

func funcName(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturn(a string, b string) (string, string) {
	return a, b
}

func variadictFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(funcName(10))

	fmt.Println(namedReturn("Robson"))

	fmt.Println(moreReturn("Robson", "Silva"))

	fmt.Println(variadictFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15))

	fmt.Println("################################")

	z := 0
	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
	fmt.Println(add())

	fmt.Println(funcInsideFunc()())

}
