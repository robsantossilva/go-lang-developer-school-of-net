package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

type SuoerCar struct {
	Car
	CanFly bool
}

func (sc SuoerCar) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s\nCanFly: %t", sc.Name, sc.Year, sc.Color, sc.CanFly)
}

func main() {

	car1 := Car{"Corolla", 2017, "White"}
	fmt.Println(car1.info())

	fmt.Println("#############################")

	sCar := SuoerCar{
		Car: Car{
			"Fusca",
			2030,
			"Blue",
		},
		CanFly: true,
	}

	fmt.Println(sCar.info())

}
