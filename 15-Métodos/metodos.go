package main

import "fmt"

type usuario struct {
	nome string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("salvando dados do %s no banco de dados\n", u.nome)
}

func(u usuario) maiorDeIdade() bool {
	return u.idade >=18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"thiago", 16}
	usuario1.salvar()

	usuario2 := usuario{"davi", 17}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
