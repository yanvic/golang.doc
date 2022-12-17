package main

import "fmt"

func main() {
	var array1 [5]string
	array1[0] = "posicao 0"
	fmt.Println(array1)

	array2 := [5]string{"numero 1", "numero 2", "numero 3", "numero 4", "numero 5"}
	fmt.Println(array2)

	array3 := [...]int{1,2,2,3,4}
	fmt.Println(array3)

	slice := []int{2,4,5,}
	slice =append(slice, 19)
	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	//arrays interno


}