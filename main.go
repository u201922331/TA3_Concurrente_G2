package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Dado() int {
	d1 := rand.Intn(6) + 1                        // Dado 1
	d2 := rand.Intn(6) + 1                        // Dado 2
	s := int(math.Pow(-1, float64(rand.Intn(2)))) // Dado Signo
	// fmt.Printf("Operacion: %d + %d * %d = ", d1, s, d2)
	return d1 + s*d2
}

func GenTablero(casillas uint) []rune {
	tablero := make([]rune, casillas)
	for i := range tablero {
		tablero[i] = '.' // . -> Casillas en blanco
	}
	tablero[0] = '#'              // # -> Inicio
	tablero[len(tablero)-1] = '$' // $ -> Fin

	umbral := int(float64(casillas) * 0.4) // Solo se llenará el 40% de las casillas en blanco con casillas especiales

	i := 0
	for i < umbral {
		idx := rand.Intn(int(casillas))
		// No afectar las casillas inicial y final
		if idx == 0 || idx == len(tablero)-1 {
			continue
		}

		// 1 -> +3 espacios
		// 2 -> -3 espacios
		// 3 -> REGRESA AL PRINCIPIO
		switch rand.Intn(3) + 1 {
		case 1:
			tablero[idx] = '1'
		case 2:
			tablero[idx] = '2'
		case 3:
			tablero[idx] = '3'
		}
		i++
	}

	return tablero
}

func main() {
	tablero := GenTablero(15)
	fmt.Printf("%c\n", tablero)
	for i := 0; i < 5; i++ {
		fmt.Println(Dado())
	}
}
