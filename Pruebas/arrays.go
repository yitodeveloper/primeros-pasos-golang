package main

import (
	"fmt"
)

func mainArrays() {
	var x [4][6][2]string
	fmt.Println(x)

	y := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(y)

	y[3] = 48

	fmt.Println(y)

	j := [...]int{
		12,
		24,
		32,
	}
	fmt.Println(j)

	//imprimiendo el tamaño de un array len(y)
	fmt.Println(len(y))
}
