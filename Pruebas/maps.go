package main

import "fmt"

// Alumno de la clase
type Alumno struct {
	id     int
	nombre string
	edad   int
}

func mainMaps() {
	x := make(map[int]Alumno) // Se puede agregar la capacidad como segundo parametro

	al := Alumno{
		id:     45673,
		nombre: "Juan",
		edad:   25,
	}

	x[456789] = al

	fmt.Println(x)

	frutitas := map[int]string{
		1: "Melonosaurio",
		2: "OtraFruta",
	}

	fmt.Println(frutitas)
	delete(frutitas, 2)
	frutitas[1] += " Rico"
	fmt.Println("La frutita es : " + frutitas[1])

	// Comprobar si un vaor existe

	if _, ok := frutitas[2]; ok {
		fmt.Println("Existe la frutita")
	} else {
		fmt.Println("No existe")
	}
}
