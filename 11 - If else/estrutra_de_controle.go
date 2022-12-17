package main

import (
	"fmt"
)
//olhar a biblipoteca rand, ajeitar o nome , fazer pastas para cadas
func main() {
	numero := 22
	numeroSorteado := 20

	if numero == numeroSorteado {
		fmt.Println("voce ganhou")
	}else if numeroSorteado > 20 {
		fmt.Println("voce chegou perto")
	}else{
		fmt.Println("voce perdeu")
	}
}
