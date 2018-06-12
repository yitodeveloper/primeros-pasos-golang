package main

import (
	"fmt"
)

func imprimir() {
	fmt.Println("Hola Rodrigo")
	defer func() {
		cadena := recover()
		fmt.Println(cadena)
	}()
	panic("Todo esta mal")
}
func mainPanic() {
	imprimir()
	fmt.Println("Hola Main")
}
