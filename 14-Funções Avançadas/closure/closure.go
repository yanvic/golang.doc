package main

import "fmt"

func closure() func() {
	texto := "função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main() {
	texto := "função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
