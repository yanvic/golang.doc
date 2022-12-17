package main

import "fmt"

func main() {
//ponteiro é uma referencia de memoria
	var variavel1 int
	var ponteiro *int

	variavel1 = 300
	ponteiro = &variavel1
	fmt.Println(variavel1, ponteiro)
	fmt.Println(variavel1, *ponteiro) //desferenciacao
}
