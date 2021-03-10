package main

import "fmt"

//Fan out

/**
* Um canal que se distribui em diversos canais
 */

func main() {
	c := generate(4, 10)

	//fan out
	d1 := divide(c)
	d2 := divide(c)

	for n1 := range d1 {
		fmt.Println(n1)
	}

	for n2 := range d2 {
		fmt.Println(n2)
	}
	fmt.Println("Finished...")
}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)
	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()
	return channel
}
