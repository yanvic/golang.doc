package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "numero inválido"
	}
}

func diaSemana2(numero int) string {
	var diaSemana2 string
	switch {
	case numero == 1:
		diaSemana2 = "Domingo"
	case numero == 2:
		diaSemana2 = "Segunda"
	case numero == 3:
		diaSemana2 = "terça"
	case numero == 4:
		diaSemana2 = "Quarta"
	case numero == 5:
		diaSemana2 = "Quinta"
	case numero == 6:
		diaSemana2 = "Sexta"
	case numero == 7:
		diaSemana2 = "Sábado"
	default:
		diaSemana2 = "numero inválido"
	}
	return diaSemana2
}
//fallthrough - joga para proxima condição



func main() {
	dia := diaSemana(1)
	fmt.Println(dia)
	dia2 := diaSemana2(3)
	fmt.Println(dia2)
}