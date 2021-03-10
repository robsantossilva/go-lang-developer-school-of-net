package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("#######################")

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("#######################")

	for {
		x++

		if x == 50 {
			fmt.Println(x)
			continue
		}

		if x == 100 {
			break
		}
	}
	fmt.Println(x)

}
