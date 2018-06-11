package main

import "fmt"

func main() {
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
}
