package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go runProcess("A", 20)
	go runProcess("B", 20)

	var s string
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}
