package main

import (
	"fmt"
	"sync"
	"time"
)

type Mineral struct {
	ID   int
	Tipo string
	Peso int
}

//func generateRandomNumber(min, max int) int {
//	return rand.IntN(max-min) + min
//}

func drones(tolva chan<- Mineral, unidades int) {
	for i := 0; i < unidades; i++ {
		time.Sleep(time.Millisecond * 150)
		tolva <- Mineral{
			ID:   i,
			Tipo: "Hierro",
			Peso: 5,
		}
		fmt.Printf("🚀 Mineral [%d] enviado a la tolva\n", i)
	}
	close(tolva)
	fmt.Println("Drones terminaron")
}

func refineria(tolva <-chan Mineral, wg *sync.WaitGroup) {
	defer wg.Done()
	for mineral := range tolva {
		fmt.Printf("🔥 Procesando mineral [%d]...\n", mineral.ID)
		time.Sleep(time.Second)
		fmt.Printf("✅ Mineral [%d] fundido\n", mineral.ID)
	}
	fmt.Println("Refinería termino")
}

func main() {
	tolva := make(chan Mineral, 5)
	unidadesExtraccion := 15

	var wg sync.WaitGroup
	wg.Add(1)

	go drones(tolva, unidadesExtraccion)
	go refineria(tolva, &wg)

	wg.Wait()
	fmt.Println("Programa finalizado\n")

}
