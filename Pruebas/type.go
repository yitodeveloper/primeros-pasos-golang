package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Errores ...
var (
	ErrorSinCaracteristicas = errors.New("El articulo no tiene caracteristicas")
	ErrorNoInicializado     = errors.New("No se ha contruido el mapa de caracteristicas")
	ErrorPorDefecto         = errors.New("Ha ocurrido un error")
)

// Articulo ...
type Articulo struct {
	Nombre string
	Categoria
	Caracteristicas map[string]Caracteristica
	Precio          int
}

// AgregarCaracteristica ...
func (a *Articulo) AgregarCaracteristica(nombre, valor string) Articulo {
	if a.Caracteristicas == nil {
		a.Caracteristicas = make(map[string]Caracteristica)
	}
	var c Caracteristica
	c.Valor = valor
	a.Caracteristicas[nombre] = c

	return *a
}

//CheckCaracteristicas ...
func (a *Articulo) CheckCaracteristicas() error {
	if a.Caracteristicas == nil {
		return ErrorNoInicializado
	}

	if len(a.Caracteristicas) == 0 {
		return ErrorSinCaracteristicas
	}

	return nil
}

// Categoria ...
type Categoria struct {
	Codigo string
	Nombre string
}

// Caracteristica ...
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

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
func main() {
	var a Articulo
	a.Nombre = "Mouse Inalambrico"
	a.Categoria.Nombre = "Computadores"
	a.Codigo = "FAD345"
	a.Precio = 500
	err := a.CheckCaracteristicas()
	checkError(err)
	a.AgregarCaracteristica("color", "naranjo")
	a.AgregarCaracteristica("tamaño", "20x10x5 cm")
	fmt.Println(a)
}
