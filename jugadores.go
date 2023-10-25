package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Jugador struct {
    Nombre  string
    Fichas  []int
}

func (j *Jugador) LanzarDado() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(6) + 1
}

func jugadorHaGanado(jugador *Jugador) bool {
	
    //return true
}

func main() {
    jugadores := []Jugador{
        {"Jugador 1", []int{0, 0, 0, 0}},
        {"Jugador 2", []int{0, 0, 0, 0}},
        {"Jugador 3", []int{0, 0, 0, 0}},
        {"Jugador 4", []int{0, 0, 0, 0}},
    }

    jugadorActual := 0

    for {
        jugador := &jugadores[jugadorActual]

        dado := jugador.LanzarDado()
        fmt.Printf("%s lanzó un %d\n", jugador.Nombre, dado)
        if jugadorHaGanado(jugador) {
            fmt.Printf("¡El jugador %s ha ganado!\n", jugador.Nombre)
            break
        }

        jugadorActual = (jugadorActual + 1) % len(jugadores)
    }
}


