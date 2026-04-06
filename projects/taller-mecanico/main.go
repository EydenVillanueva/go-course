package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Vehiculo struct {
	ID              int
	Tipo            string
	CostoReparacion int
}

var totalVehiculos int
var nMecanicos int
var caja int
var nCajeros int
var tiposDeVehiculo []string
var elevadores int
var limiter chan struct{}

func init() {
	tiposDeVehiculo = []string{"Auto", "Camioneta"}
	totalVehiculos = 20
	nMecanicos = 5
	caja = 0
	elevadores = 2
	nCajeros = 3
	limiter = make(chan struct{}, elevadores)
}

func generadorDeTipo() Vehiculo {
	switch esCamioneta := rand.IntN(100) < 70; {
	case esCamioneta == true:
		return Vehiculo{ID: 0, Tipo: tiposDeVehiculo[1], CostoReparacion: 200}
	default:
		return Vehiculo{ID: 0, Tipo: tiposDeVehiculo[0], CostoReparacion: 100}
	}
}

func generadorDeVehiculos(chRecepcion chan<- Vehiculo) {
	defer close(chRecepcion)
	for i := 0; i < totalVehiculos; i++ {
		tipo := generadorDeTipo()
		tipo.ID = i + 1

		fmt.Printf("- Vehiculo %d %s ha llegado\n", tipo.ID, tipo.Tipo)
		chRecepcion <- tipo
		time.Sleep(time.Millisecond * 200)
	}
}

func mecanico(chRecepcion <-chan Vehiculo, chCaja chan<- int, wgReparacion *sync.WaitGroup) {
	defer func() {
		wgReparacion.Done()
	}()

	for vehiculo := range chRecepcion {
		limiter <- struct{}{}
		fmt.Printf("Reparando vehiculo[%d]%s\n", vehiculo.ID, vehiculo.Tipo)
		var tiempoReparacion time.Duration

		switch esCamioneta := vehiculo.Tipo == "Camioneta"; {
		case esCamioneta:
			tiempoReparacion = time.Millisecond * time.Duration(3000)
		default:
			tiempoReparacion = time.Millisecond * time.Duration(1000)
		}
		time.Sleep(tiempoReparacion)
		fmt.Printf("[%d]%s Reparacion completada\n", vehiculo.ID, vehiculo.Tipo)
		<-limiter
		chCaja <- vehiculo.CostoReparacion
	}
}

func cobrarReparaciones(chCaja <-chan int, wgCaja *sync.WaitGroup, mu *sync.Mutex) {
	defer func() {
		wgCaja.Done()
	}()

	for coste := range chCaja {
		mu.Lock()
		caja += coste
		mu.Unlock()
	}
}

func main() {
	chRecepcion := make(chan Vehiculo, totalVehiculos)
	chCaja := make(chan int)

	var wgReparacion sync.WaitGroup // wg para esperar a los mecanicos reparar
	var wgCaja sync.WaitGroup       // wg para esperar a la caja cobrar
	var mu sync.Mutex

	go generadorDeVehiculos(chRecepcion)

	for i := 0; i < nMecanicos; i++ {
		wgReparacion.Add(1)
		go mecanico(chRecepcion, chCaja, &wgReparacion)
	}

	go func() {
		wgReparacion.Wait()
		close(chCaja)
	}()

	for i := 0; i < nCajeros; i++ {
		wgCaja.Add(1)
		go cobrarReparaciones(chCaja, &wgCaja, &mu)
	}

	wgCaja.Wait()

	fmt.Printf("\nTotal en la caja después de las reparaciones:\n %d\n", caja)

}
