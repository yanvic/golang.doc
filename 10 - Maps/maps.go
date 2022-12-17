package main

import "fmt"

func main() {
	//[tipos das chaves] fora é o tipo dos valores
	usuario := map[string]string{
		"nome": "victor",
		"sobrenome": "silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "yan",
			"sobrenome": "victor",
		},
		"curso": {
			"nome": "direito",
			"modalidade": "ead",
		},
	}

	fmt.Println(usuario2)
}