package main

import "fmt"

var n int

func main() {
	fmt.Println("função main")
	println(n)
}
//serve para inicializar
func init()  {
	fmt.Println("função init")
	n = 5
}
