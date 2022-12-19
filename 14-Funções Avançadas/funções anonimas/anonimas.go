package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("numeros sorteado -> %d %s", 8, texto)
	}("funcao anonima")
	fmt.Println(retorno)
}
