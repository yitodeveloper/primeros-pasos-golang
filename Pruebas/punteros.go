package main

import (
	"fmt"
)

func incrementar(numero *int) {
	*numero++
}

func mainPuntero() {
	b := 12
	fmt.Println("El valor de b es ", b)
	fmt.Println("La dirección de b es ", &b)

	a := &b

	fmt.Println(a)
	fmt.Println(*a)

	*a = 34
	fmt.Println("El valor de b es ", b)
	b++
	fmt.Println("El valor de *a es ", *a)

	d := new(int)
	fmt.Println("El valor de d es ", d)
	fmt.Println("La dirección de d es ", *d)

	numero := 1
	incrementar(&numero)
	fmt.Println("El valor del numero es ", numero)
	incrementar(&numero)
	fmt.Println("El valor del numero es ", numero)
	incrementar(&numero)
	fmt.Println("El valor del numero es ", numero)
}
