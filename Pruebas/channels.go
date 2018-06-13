package main

import (
	"fmt"
	"strconv"
	"time"
)

var numeroPublico int

func imprimirPing(c chan string) {
	var contador int
	for {
		contador++
		time.Sleep(time.Second * 2)
		x, ok := <-c
		if !ok {
			break
		}

		fmt.Println(x, " - Cliente>", contador, " Publico Cliente", numeroPublico)

	}
}

func enviarPing(c chan string) {
	var contador int
	for i := 0; i < 10; i++ {

		contador++
		fmt.Println("Enviando data...")
		c <- "Servidor>" + strconv.Itoa(contador) + " Publico Cliente" + strconv.Itoa(numeroPublico)

		//time.Sleep(time.Second * 2)
	}
	close(c)
}

func mainChannel() {
	/*
		Para los channel se puede determinar la direccionalidad (undirectional channel)
		en la declaracion del channel por Ejemplo
		chan<-    Channel de solo salida
		<-chan    Channel de solo entrada

	*/
	channel := make(chan string)
	/*
		Buffered Channel make(chan string,4) Encola mensajes enviados antes de blockearse
		Unbuffered Channel make(chan string) Se blockea por cada mensaje enviado
	*/
	go enviarPing(channel)
	go imprimirPing(channel)

	/*for {
		numeroPublico++
		println("Num : ", numeroPublico)
		time.Sleep(time.Second * 2)
	}*/
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin del Programa")
}
