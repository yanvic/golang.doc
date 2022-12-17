package main
import "fmt"
//import "time"

func main() {
//	i := 0
//	for i < 10 {
//		i++
//		fmt.Println("contando")
//		time.Sleep(time.Second)
//	}

//	for a := 0; a < 30; a+= 5 {
//		fmt.Println("contando", a)
//		time.Sleep(time.Second)
//	}
	nomes := [3]string{"joao", "thiago", "pedro"}

	for posicao, nome := range nomes {
		fmt.Println(posicao, nome)
	}
	for indice, letra := range "abcdefg" {
		fmt.Println(indice, string(letra))
	}
	usuario := map[string]string{
		"nome": "leo",
		"sobrenome": "pereira",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}