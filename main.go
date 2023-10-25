package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Dice() int {
	a := rand.Intn(6) + 1                         // Dado 1
	b := rand.Intn(6) + 1                         // Dado 2
	s := int(math.Pow(-1, float64(rand.Intn(2)))) // Dado Signo
	fmt.Printf("Operacion: %d + %d * %d = ", a, s, b)
	return a + s*b
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(Dice())
	}
}
