package main

import "fmt"

var b int = 1

const (
	aa        = 55
	bb        = 66
	cc        = 77
	dd string = "Oiiii"
	Ee int    = 123 //Constant visivel fora do pacote
)

const xvz int = 1333

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Uooouu`

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

	//CONSTANTS
	const xpto = 10
}
