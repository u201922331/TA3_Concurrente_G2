package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	Nro_casillas = 72
)

type Jugador struct {
	Nombre             string
	Fichas             int
	MovimientosTotales int
}

type Tablero struct {
	jugadores []Jugador
	tablero   []rune
}

func (t Tablero) MostrarJugadores(j []Jugador) {

	for i := 0; i < 4; i++ {
		nuevo := make([]rune, len(t.tablero))
		copy(nuevo, t.tablero)
		if j[i].MovimientosTotales <= 0 {
			nuevo[0] = 'X'
		}
		if j[i].MovimientosTotales > len(t.tablero)-1 {
			nuevo[len(t.tablero)-1] = 'X'
		} else {
			nuevo[j[i].MovimientosTotales] = 'X'
		}

		fmt.Printf("%c\n", nuevo)
	}
}

func Dado() int {
	d1 := rand.Intn(6) + 1                        // Dado 1
	d2 := rand.Intn(6) + 1                        // Dado 2
	s := int(math.Pow(-1, float64(rand.Intn(2)))) // Dado Signo
	fmt.Printf("Operacion: %d + (%d * %d)", d1, s, d2)
	return d1 + s*d2
}

func GenTablero(casillas int) []rune {
	tablero := make([]rune, casillas)
	for i := range tablero {
		tablero[i] = '_' // . -> Casillas en blanco
	}
	tablero[0] = '#'              // # -> Inicio
	tablero[len(tablero)-1] = '#' // $ -> Fin

	umbral := int(float64(casillas) * 0.4) // Solo se llenará el 40% de las casillas en blanco con casillas especiales

	for i := 0; i < umbral; i++ {
		idx := rand.Intn(int(casillas))
		// No afectar las casillas inicial y final
		if idx == 0 || idx == len(tablero)-1 {
			continue
		}

		/*
			1 -> +3 espacios
			2 -> -3 espacios
			3 -> REGRESA AL PRINCIPIO
		*/
		switch rand.Intn(3) + 1 {
		case 1:
			tablero[idx] = '1'
		case 2:
			tablero[idx] = '2'
		case 3:
			tablero[idx] = '3'
		}
	}
	return tablero
}

func Meta(j *Jugador) bool {
	if j.MovimientosTotales >= Nro_casillas-1 {
		return true
	}
	return false
}

func main() {
	jugadores := []Jugador{
		{"J1", 0, 0},
		{"J2", 0, 0},
		{"J3", 0, 0},
		{"J4", 0, 0},
	}
	t := Tablero{jugadores, GenTablero(Nro_casillas)}

	jugadorActual := 0

	for {
		if jugadorActual == 0 {
			fmt.Printf("----------------------------------------------\n")
		}

		j := &jugadores[jugadorActual]

		dado := Dado()
		j.MovimientosTotales = j.MovimientosTotales + dado

		if j.MovimientosTotales < 0 {
			j.MovimientosTotales = 0
		}
		if j.MovimientosTotales > 71 {
			j.MovimientosTotales = 71
		}

		if t.tablero[j.MovimientosTotales] == '1' {
			fmt.Printf("\t\tCASILLA ESPECIAL: +3 ")
			j.MovimientosTotales = j.MovimientosTotales + 3
			if j.MovimientosTotales > 71 {
				j.MovimientosTotales = 71
			}
		}
		if t.tablero[j.MovimientosTotales] == '2' {
			fmt.Printf("\t\tCASILLA ESPECIAL: -3 ")
			j.MovimientosTotales = j.MovimientosTotales - 3
			if j.MovimientosTotales < 0 {
				j.MovimientosTotales = 0
			}
		}
		if t.tablero[j.MovimientosTotales] == '3' {
			fmt.Printf("\t\tCASILLA ESPECIAL: 0 ")
			j.MovimientosTotales = 0
		}

		fmt.Printf("\t\t%s OBTUVO: %d, SUM: %d", j.Nombre, dado, j.MovimientosTotales)

		if Meta(j) {
			j.Fichas++
			if j.Fichas == 4 {
				fmt.Printf("\t\t¡%s GANO!, cantidad de fichas metidas: %d\n", j.Nombre, j.Fichas)
				break
			}

			fmt.Printf("\t\t¡%s metio una ficha!, cantidad de fichas metidas: %d", j.Nombre, j.Fichas)
			j.MovimientosTotales = 0
		}

		jugadorActual = (jugadorActual + 1) % len(jugadores)

		fmt.Printf("\n")
	}

	t.MostrarJugadores(jugadores)
}
