package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//Waiting Grouos

var waitGroup sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println(runtime.NumCPU())
	waitGroup.Add(2)
	go runProcess("A", 20)
	go runProcess("B", 20)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}
