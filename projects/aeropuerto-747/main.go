package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Maleta struct {
	ID           int
	Peso         int
	EsSospechosa bool
}

var nGuardias int
var limiter chan struct{}
var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	nGuardias = 5
	limiter = make(chan struct{}, 2) // 2 maquinas de rayos x
}

func randomInt(min, max int) int {
	return rand.IntN(max-min) + min
}
func generadorDeSospecha() bool {
	return rand.IntN(100) < 10
}

// genera las maletas y las pone en la cinta de check-in
func pasajeros(chCheckin chan<- Maleta) {
	defer close(chCheckin)
	nMaletas := randomInt(10, 15)
	fmt.Printf("---------- Numero de Maletas: %d ----------\n\n", nMaletas)

	for i := 0; i < nMaletas; i++ {
		chCheckin <- Maleta{
			ID:           i + 1,
			Peso:         randomInt(2, 5), //kg
			EsSospechosa: generadorDeSospecha(),
		}
		time.Sleep(time.Millisecond * 200)
	}
}

// Voy a sacar maletas del chCheckin
// Voy a mandar las maletas ya revisadas a chCargo para su carga al avión
func guardiasDeSeguridad(chCheckin <-chan Maleta, wgCheckIn *sync.WaitGroup, chCarga chan<- Maleta) {
	defer func() {
		wgCheckIn.Done()
	}()

	for maleta := range chCheckin {
		limiter <- struct{}{}

		fmt.Printf("Revisando Maleta [%d]\n", maleta.ID)
		revision := time.Millisecond * time.Duration(maleta.Peso*100)

		if !maleta.EsSospechosa {
			time.Sleep(revision) // 2 kg = 200ms, 5kg = 500ms
		} else {
			fmt.Printf("Maleta sospechosa, revisión manual [%d]\n", maleta.ID)
			time.Sleep(revision * 2) // El doble de tiempo para una revisión manual
		}
		fmt.Printf("Revisión de Maleta: [%d] Terminada enviando a la cinta de carga\n", maleta.ID)
		chCarga <- maleta
		<-limiter
	}
}

func cargarAvion(chCarga <-chan Maleta, wgCargo *sync.WaitGroup) {
	defer wgCargo.Done()

	for maletaRevisada := range chCarga {
		fmt.Printf("Maleta[%d] Cargada!\n", maletaRevisada.ID)
	}
}

func main() {
	chCheckIn := make(chan Maleta, 10) // cinta check-int
	chCarga := make(chan Maleta, 10)   // cinta de carga

	var wgCheckIn sync.WaitGroup // wg para esperar los guardias de seguridad
	var wgCarga sync.WaitGroup   // wg para esperar la carga de maletas al avión

	go pasajeros(chCheckIn)

	// Proceso de revision de las maletas
	for i := 0; i < nGuardias; i++ {
		wgCheckIn.Add(1)
		go guardiasDeSeguridad(chCheckIn, &wgCheckIn, chCarga)
	}

	go func() {
		wgCheckIn.Wait()
		close(chCarga)
	}()

	wgCarga.Add(1)
	go cargarAvion(chCarga, &wgCarga)
	wgCarga.Wait()
}
