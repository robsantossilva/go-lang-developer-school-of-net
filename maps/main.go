package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	fmt.Println("############################")

	delete(m, "c")
	fmt.Println(m)

	fmt.Println("############################")

	_, exists := m["c"]
	fmt.Println(exists)

	fmt.Println("############################")

	value, exists := m["c"]
	fmt.Println(value)

	fmt.Println("############################")

	var x = map[string]int{}
	fmt.Println(x)

	fmt.Println("############################")

	newMap := map[string]int{"x": 5, "y": 10}
	fmt.Println(newMap)

	fmt.Println("############################")

	if value, exists := m["a"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("Ops!")
	}
}
