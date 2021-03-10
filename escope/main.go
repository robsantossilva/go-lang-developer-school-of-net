//go run ./escope/*.go
package scope

import "fmt"

var Y int = 20

// func main() {
// 	x := 10
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// 	printZ()
// }

func printX() {
	//fmt.Println(x) //x esta no escopo de main()
	//fmt.Println(y) //y esta no escopo global
}

func PrintY() {
	fmt.Println(Y)
}
