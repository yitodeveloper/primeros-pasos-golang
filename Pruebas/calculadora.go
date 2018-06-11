package main

import (
	"fmt"
	"os"
	"strconv"
)

const welcomeSplash = `
	***********************************************************************
	***                                                                 ***
	***            Bienvenido a la Calculadora                          ***
	***                           Powered by @yitodeveloper             ***
	***                                                                 ***
	***********************************************************************
`
const menu = `
	Ingrese una de las siguientes opciones para realizar la operación deseada.
	1) Suma
	2) Resta
	3) Multiplicación
	x) Salir de la aplicación
`

func print(text string) {
	if text != "" {
		fmt.Fprintln(os.Stdout, text)
	}
}

func main() {
	print(welcomeSplash)

	salir := true
	var opcion string

	for salir {
		print(menu)
		fmt.Scanln(&opcion)
		if opcion == "1" {
			var total int
			var cantidad int
			print("Ingrese la cantidad de numeros que desea contar")
			fmt.Scanln(&cantidad)
			for i := 1; i <= cantidad; i++ {
				var numero int
				print("Ingrese un numero a sumar - n°" + strconv.Itoa(i))
				fmt.Scanln(&numero)
				total += numero
			}
			print("El total de los numeros ingresados es " + strconv.Itoa(total))
		} else if opcion == "2" {
			var total int
			var cantidad int
			print("Ingrese la cantidad de numeros que desea restar")
			fmt.Scanln(&cantidad)
			for i := 1; i <= cantidad; i++ {
				var numero int
				print("Ingrese un numero a restar - n°" + strconv.Itoa(i))
				fmt.Scanln(&numero)
				if i != 1 {
					total -= numero
				} else {
					total += numero
				}

			}
			print("El total de los numeros ingresados es " + strconv.Itoa(total))
		} else if opcion == "3" {
			var total int
			var cantidad int
			print("Ingrese la cantidad de numeros que desea multiplicar")
			fmt.Scanln(&cantidad)
			total = 1
			for i := 1; i <= cantidad; i++ {
				var numero int
				print("Ingrese un numero a multiplicar - n°" + strconv.Itoa(i))
				fmt.Scanln(&numero)
				total *= numero

			}
			print("El total de los numeros ingresados es " + strconv.Itoa(total))
		} else if opcion == "x" {
			salir = false
			print("Saliendo del programa, Hasta luego :)")
		} else {
			print("Digite una de las opciones validas")
		}

	}

}
