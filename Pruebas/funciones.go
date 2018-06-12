package main

import (
	"fmt"
	"os"
)

//Variadic function
func sumarRestar(numeros ...int) (r int, j int) {
	for key, num := range numeros {
		r += num
		if key != 1 {
			j -= num
		} else {
			j += num
		}

	}
	return
}

func decirHola(nombre string) {
	fmt.Println("Hola " + nombre + "!")
}

func primera() {
	fmt.Println("Soy la primera")
}

func segunda() {
	fmt.Println("Soy la segunda")
}

func mainFunciones() {
	fmt.Println(sumarRestar())

	numeros := []int{
		24,
		12,
		1,
	}

	fmt.Println(sumarRestar(numeros...))

	//Closure
	saludar := decirHola

	saludar("Juan")

	saludar2 := func() {
		fmt.Println("Saludo nuevamente")
	}

	saludar2()

	//Las funciones tienen una firma que determina el typo de la funcion, si la funcion recibe distintos parametros
	// y devuelve distintos valores la firma es distinta, y cuando esta asignada una funcion a una variable y esta
	// se reasigna con una de distinta firma, el compilador de go entrega error

	//Funciones anonimas

	//Donde se requiera una función como parametro la función se puede escribir inmediatamente dentro de la declaración como parametro

	// La propiedad defer
	defer primera()
	segunda()

	file, err := os.Open("loremd.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]byte, 200)
	c, err := file.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cantidad de byte leidos : %d\n Texto leido: %q\n Error: %v\n", c, data, err)

}
