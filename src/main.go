package main

import "fmt"

func main() {

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("Este es el valor de pi:", pi)
	fmt.Println("Este es el valor de pi2:", pi2)

	// Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println("[Base, Altura, Area]: [", base, ", ", altura, ", ", area, "]")

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado:", areaCuadrado)
}
