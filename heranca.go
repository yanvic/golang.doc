package main

import "fmt"

type pessoa struct {
	nome string
	idade int
	altura float64
	peso float64
}
type estudante struct {
	pessoa
	curso string
	modalidade string
	faculdade string
}
func main() {
	p1 := pessoa{"jonas", 20, 1.85, 87.45}
	fmt.Println(p1)
	e1 := estudante{p1, "ads", "noturno", "uniateneu"}
	fmt.Println(e1)
}
