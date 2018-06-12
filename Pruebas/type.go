package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Articulo struct {
	Nombre string
	Categoria
	Caracteristicas map[string]Caracteristica
	Precio          int
}

func (a *Articulo) AgregarCaracteristica(nombre, valor string) Articulo {
	if a.Caracteristicas == nil {
		a.Caracteristicas = make(map[string]Caracteristica)
	}
	var c Caracteristica
	c.Valor = valor
	a.Caracteristicas[nombre] = c

	return *a
}

type Categoria struct {
	Codigo string
	Nombre string
}

type Caracteristica struct {
	Nombre string
	Valor  string
}

func (a Articulo) String() string {
	car := "\n Caracteristicas : "
	for n, c := range a.Caracteristicas {
		car += "\n            " + strings.Title(n) + ": " + c.Valor
	}
	return fmt.Sprintf("[" + strings.ToUpper(a.Categoria.Nombre) + "] " + a.Nombre + " $" + strconv.Itoa(a.Precio) + car)
}

func main() {
	var a Articulo
	a.Nombre = "Mouse Inalambrico"
	a.Categoria.Nombre = "Computadores"
	a.Codigo = "FAD345"
	a.Precio = 500
	a.AgregarCaracteristica("color", "naranjo")
	a.AgregarCaracteristica("tamaño", "20x10x5 cm")
	fmt.Println(a)
}
