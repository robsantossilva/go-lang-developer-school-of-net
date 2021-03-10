package main

import (
	"fmt"
	"package/car"
)

func main() {

	car := car.Car{"Gol", "Black"}

	fmt.Println(car.Start())

}
