package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"car_name"`
	Year  int    `json:"car_year"`
	Color string `json:"car_color"`
}

func main() {
	var car Car
	data := []byte(`{"car_name":"Gol","car_year":2017,"car_color":"Black"}`)

	json.Unmarshal(data, &car)

	fmt.Println(car.Name, car.Year, car.Color)
}
