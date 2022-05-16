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

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División", result)

	// Módulo(resto)
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decrementa", x)
}
