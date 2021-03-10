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
	car := Car{"Gol", 2017, "Yellow"}
	result, _ := json.Marshal(car)
	fmt.Println(string(result))
}
