package main

import (
	"fmt"
)

func mainSlices() {
	var j []int
	fmt.Println(j)
	x := []string{
		"Manzana",
		"Pera",
		"Durazno",
		"Piña",
	}
	fmt.Println(x)

	// Construir un slice nuevo sin valores
	/*y := make([]int, 6)
	fmt.Println(y)*/

	x = append(x, "Sandia", "platano", "kiwi")

	fmt.Println(x)

	frutas := []string{
		"Melon",
	}

	copy(x, frutas)
	fmt.Println(frutas)
	fmt.Println(x)

	for _, v := range x {
		fmt.Println("La fruta mas rica del mundo es " + v)
	}

	cadena := "Hola a todos me gusta programar"

	for _, letra := range cadena {
		fmt.Printf("La letra es la siguiente %q \n", letra)
	}
}
