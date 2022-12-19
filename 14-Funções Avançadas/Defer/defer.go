package main

import "fmt"
//defer = adiar até o ultimo momento
func main() {
	fmt.Println("entrando no sistema")
	fmt.Println(alunoAprovado(5.5,8))
	fmt.Println(alunoAprovado2(4.5,8))
}

func alunoAprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	if media >= 6 {
		fmt.Println("media sendo calculada")
		return true
	}
	fmt.Println("media sendo calculada")
	return false
}

func alunoAprovado2(n1, n2 float32) bool {
	defer println("media sendo calculada, resultado sera retornado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}
