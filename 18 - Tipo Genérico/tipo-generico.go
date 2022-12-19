package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}] interface{}{
		1: true,
		"string": 3,
		true: int(8),
	}
	fmt.Println(mapa)
}

