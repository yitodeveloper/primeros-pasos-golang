package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func lanzarRutina(nombre string) {
	fmt.Println("Inicializando hilo [" + nombre + "]")
	defer wg.Done()
	for contador := 1; contador <= 20; contador++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println(nombre + "> Contador en " + strconv.Itoa(contador))
	}
}

func mainOld() {
	wg.Add(2)
	fmt.Println("Comenzando las gorutinas")
	go lanzarRutina("Queso")
	go lanzarRutina("Palta")
	fmt.Println("Esperando a finalizar")
	wg.Wait()
	fmt.Println("Finalizo el programa")
}

func mainFib() {
	go animacion(100 * time.Millisecond)
	const n = 45
	resultado := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, resultado)
}

func animacion(retraso time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
