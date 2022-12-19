package main

import "fmt"

func recuperandoSistema()  {
	if r := recover(); r != nil {
		fmt.Println("sistema recuperado com sucesso!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperandoSistema()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("a media é exatamente 6!")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("encerrando")
}
// ele acaba com o programa