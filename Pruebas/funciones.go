package main

import "fmt"

func sumar(x, y int) (r int) {
	r = x + y
	return
}

func main() {
	fmt.Println(sumar(2, 3))
}
