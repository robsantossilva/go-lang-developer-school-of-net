package main

import "fmt"

func main() {
	slice := make([]int, 5)
	fmt.Println(slice)

	slice = append(slice, 1)
	fmt.Println(slice)

	fmt.Println(len(slice))

	fmt.Println("#######################################")

	/** Capacity - When it exceeds capacity, it is automatically folded*/
	sliceCapacity := make([]int, 5, 6)
	fmt.Println(sliceCapacity)
	fmt.Println(len(sliceCapacity))
	fmt.Println(cap(sliceCapacity))
	fmt.Println("-----------------")
	sliceCapacity = append(sliceCapacity, 12)
	fmt.Println(sliceCapacity)
	fmt.Println(len(sliceCapacity))
	fmt.Println(cap(sliceCapacity))
	fmt.Println("-----------------")
	sliceCapacity = append(sliceCapacity, 24, 32)
	fmt.Println(sliceCapacity)
	fmt.Println(len(sliceCapacity))
	fmt.Println(cap(sliceCapacity))

	fmt.Println("#######################################")

	slice2 := make([]int, 2, 5)
	slice3 := slice2
	slice2[0] = 10 //Modify slice2 and slice3
	// just modify slice2 and return a new slice because the capacity has burst
	// whenever a slice overflows a new array is created
	slice2 = append(slice2, 1, 2, 3, 4)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("-----------------")
	slice2[1] = 333
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("#######################################")

	//create a slice without make
	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Better 2",
	}

	fmt.Println(sliceString[0])
	fmt.Println(sliceString[:2])
	fmt.Println(sliceString[1:2])
	fmt.Println(sliceString[2:4])
	fmt.Println(sliceString[2:])

}
