package main

import (
	"fmt"	
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.1

	area := PI * math.Pow(raio, 2)

	fmt.Println("Área da circunferência é", area)

	const (
		c = 3
		d = 5
	)

	fmt.Println("Constantes:", c, d)

	var (
		a = 1
		b = 2
	)

	fmt.Println("Variável:", a, b)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Opa!"
	
	fmt.Println(g, h, i)

}