### Tarea Académica 3

# Informe de implementación del juego Ludo modificado
#### Integrantes
- Alumno 1
- Alumno 2
- Alumno 3

## 1. Descripción
Este informe presenta la implementación del juego Ludo modificado realizado con el lenguaje de programación GO. El juego se desarrolla en un tablero que contienen obstáculos y casillas especiales. El objetivo es que los jugadores muevan sus fichas desde el inicio hasta la meta. El juego fue implementado utilizando conceptos de concurrencia y comunicación entre procesos(canales). 

## 2. Estructura del juego 
El juego se desarrolla con las siguientes estructuras y funciones clave
- **Jugador** : Representa a un jugador con un nombre, la cantidad de fichas metidas en la meta y la cantidad de movimientos totales realizados.
- **Tablero** : Representa el tablero del juego con jugadores y casillas. El tablero contiene un slice de rune que representa las casillas del juego.
- **Función "MostrarJugadores"** : Muestra el estado de los jugadores en el tablero, indicando su posición en las casillas.
- **Función "Dado"** : Simula el lanzamiento de dados. Los jugadores lanzan dos dados, uno con valores del 1 al 6 y otro con la operación de suma o resta. Los resultados determinan cuántos pasos pueden avanzar o retroceder.
- **Función "GenTablero"** : Genera el tablero del juego con casillas en blanco, casillas de inicio, casillas de meta y casillas especiales (con obstáculos y movimientos adicionales).
- **Función "Meta"** : Verifica si un jugador ha llegado a la casilla de meta (casilla final).
- **Función "obtenerEstadoJuego"** : Obtiene el estado del juego, incluyendo los movimientos totales y fichas metidas de todos los jugadores.
- **Goroutine para Mostrar el Estado del Juego** : Se crea una goroutine para mostrar el estado del juego en tiempo real. El estado se envía a través de un canal llamado estadoJuegoChan.
  
## 3. Desarrollo del juego
- Se inicializan los jugadores con sus nombres, sin fichas metidas y movimientos totales en cero. El tablero se crea y se llena con casillas especiales y obstáculos.
- Los jugadores se turnan para lanzar los dados y avanzar en el tablero. Los resultados de los dados determinan la cantidad de pasos y la operación de suma o resta.
- Los jugadores pueden encontrar casillas especiales que les permiten avanzar o retroceder en el tablero. Si llegan a la casilla de meta, meten una ficha y reinician su posición en el tablero.
- El juego continúa hasta que un jugador haya metido cuatro fichas, lo que indica que ha ganado.

## 4. Aplicación de canales
- Se define un canal que envía y recibe información sobre el estado del juego. Este canal transmite el estado del juego a medida que los jugadores realizan sus movimientos
```go
// Define un canal para comunicar el estado del juego
estadoJuegoChan := make(chan string)
```

- Se inicia un gourotine que escucha continuamente a través del canal “estadoJuegoChan”. Este gourotine se encarga de mostrar el estado del juego en tiempo real a medida que se actualiza. Cada vez que se envía un nuevo estado del juego a través del canal, el goroutine lo captura y lo muestra en la consola.
```go
go func() {
      for estado := range estadoJuegoChan {
            fmt.Printf("Estado del juego:\n%s\n", estado)
        }
    }()
```

- Después de calcular el estado del juego, este se envía a través del canal. La función “obtenerEstadoJuego” genera una cadena que representa el estado de todos los jugadores, y esta cadena se envía al canal. El goroutine que escucha en el canal captura este estado y lo muestra en la consola.
```go
// Después de calcular el estado del juego, envía el estado al canal
        estadoJuego := obtenerEstadoJuego(jugadores)
        estadoJuegoChan <- estadoJuego
```

## 5. Conclusiones
En el proyecto desarrollado, se utilizan canales para lograr la comunicación en tiempo real del estado del juego entre los goroutines que manejan los movimientos de los jugadores y los goroutines que muestra el estado del juego. A continuación, se resumen los aspectos clave del uso de canales en el desarrollo del juego:
- Se crea un canal llamado estadoJuegoChan utilizando la función make. Este canal se utiliza para enviar y recibir información sobre el estado del juego.
- Se inicia una goroutine que escucha de manera continua en el canal estadoJuegoChan. Esta goroutine se encarga de mostrar el estado del juego en tiempo real. Cada vez que se envía un nuevo estado del juego a través del canal, la goroutine lo captura y lo muestra en la consola.
- Después de calcular el estado del juego, se envía el estado al canal mediante la expresión estadoJuegoChan <- estadoJuego. La función "obtenerEstadoJuego" genera una cadena que representa el estado de todos los jugadores, y esta cadena se envía al canal. La goroutine que escucha en el canal captura este estado y lo muestra en la consola.

Los beneficios que aporta el uso de canales en el proyecto son : 
- **Comunicación Concurrente:** Permite que las goroutines se comuniquen y compartan información de manera segura mientras se ejecutan en paralelo. Esto garantiza que la información sobre el estado del juego se actualice en tiempo real a medida que los jugadores realizan sus movimientos.
- **Sincronización:** Los canales facilitan la sincronización entre las goroutines. La goroutine encargada de mostrar el estado del juego no muestra información desactualizada, ya que espera a que se envíe un nuevo estado a través del canal.





