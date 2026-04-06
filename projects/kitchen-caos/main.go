package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Orden struct {
	ID       int
	Comida   string
	Duracion time.Duration
}

type Cliente struct {
	Nombre       string
	MaximoEspera time.Duration
	Orden
}

var menu []string
var nombres []string
var nCocineros int
var limiter chan struct{}

func init() {
	menu = []string{
		"Arrachera",
		"Pasta Alfredo",
		"Pizza de peperoni",
		"Brochetas",
		"Flan napolitano",
		"Crema de papa",
	}

	nombres = []string{
		"Paula",
		"Santiago",
		"Salma",
		"Ricardo",
		"Georgina",
		"Nancy",
	}

	nCocineros = 3

	limiter = make(chan struct{}, 2)
}

func randomInt(min, max int) int {
	return rand.IntN(max-min) + min
}

// Meseros ponen las ordenes en el channel
func generadorDeOrdenes(chOrdenes chan<- Orden, chEspera chan<- Cliente) {
	defer func() {
		close(chOrdenes)
		close(chEspera)
	}()

	nPedidos := randomInt(7, 9)

	for i := 0; i < nPedidos; i++ {
		index := randomInt(0, 5)
		cliente := Cliente{
			Nombre:       nombres[index],
			MaximoEspera: time.Duration(randomInt(1, 3)), // entre 1 y 3 tiempo de espera
			Orden: Orden{
				ID:       i + 1,
				Comida:   menu[index],
				Duracion: 0, // La duración será actualizada por los cocineros al terminar
			},
		}

		fmt.Printf("- Orden n %d (%s) registrada\n", cliente.ID, cliente.Comida)
		chOrdenes <- cliente.Orden // Registramos su orden a la cocina
		chEspera <- cliente        // Iniciamos la espera del cliente

		// Pequeña pausa para que no lleguen todos en el milisegundo 0
		time.Sleep(time.Millisecond * 200)
	}
}

// Mesa de espera
func espera(chEspera <-chan Cliente, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range chEspera {
		fmt.Printf(" Cliente (%s) inicia su espera de %d segundos!\n", c.Nombre, c.MaximoEspera)
		time.Sleep(time.Second * c.MaximoEspera)
	}
}

// Cocineros
func cocinar(chOrdenes <-chan Orden, wg *sync.WaitGroup) {
	defer wg.Done()

	for orden := range chOrdenes {
		limiter <- struct{}{}

		fmt.Printf("👨‍🍳🔥 Orden n %d (%s) en preparación !\n", orden.ID, orden.Comida)

		tiempoPorOrden := time.Duration(randomInt(1, 5)) // Entre 1 y 5 tiempo de preparación
		time.Sleep(time.Second * tiempoPorOrden)
		orden.Duracion = tiempoPorOrden

		<-limiter
	}
}

func main() {
	chOrdenes := make(chan Orden)
	chEspera := make(chan Cliente)

	var wg sync.WaitGroup

	go generadorDeOrdenes(chOrdenes, chEspera)

	for i := 0; i < nCocineros; i++ {
		wg.Add(1)
		go cocinar(chOrdenes, &wg) // Mandamos a trabajar a la cocina
	}

	wg.Add(1)
	go espera(chEspera, &wg)

	wg.Wait()
	fmt.Println("--- El restaurante ya cerró, todos a sus casas ---")
	//for {
	//	select {
	//	case c := <-chEspera:
	//		fmt.Printf(" Cliente (%s) Se fué!\n", c.Nombre)
	//	case o := <-chOrdenes:
	//		fmt.Printf("👨‍🍳🔥 Orden n %d (%s) lista en %d segundos !\n",
	//			o.ID, o.Comida, o.Duracion)
	//	case <-time.After(time.Second * 6): // Horario del restaurante se termina
	//		fmt.Println("El restaurante ya cerró, todos a sus casas")
	//		return
	//	}
	//}
}
