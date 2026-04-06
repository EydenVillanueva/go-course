package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. FABRICANTE (Productor)
// Recibe un canal "send-only" (chan<-)
func fabricMachine(nBotellas int, ch chan<- int) {
	for i := 1; i <= nBotellas; i++ {
		// Simula trabajo rápido (200ms)
		time.Sleep(time.Millisecond * 200)

		fmt.Printf("🍾 Botella [%d] FABRICADA (Intentando poner en cinta...)\n", i)

		// AQUÍ OCURRE LA MAGIA DEL BUFFER:
		// Si el buffer (3 espacios) tiene lugar, esta línea pasa inmediato.
		// Si el buffer está lleno, esta línea SE BLOQUEA y espera.
		ch <- i

		fmt.Printf("   -> Botella [%d] puesta en cinta (Buffer liberado/usado)\n", i)
	}
	// IMPORTANTE: Cerrar el canal avisa a la empacadora que no hay más.
	close(ch)
	fmt.Println("--- FABRICA CERRADA ---")
}

// 2. EMPACADORA (Consumidor)
// Recibe un canal "receive-only" (<-chan)
func empacadora(ch <-chan int, wg *sync.WaitGroup) {
	// Defer Done: Avisar al WaitGroup cuando esta función termine
	defer wg.Done()

	// range ch: Lee continuamente hasta que el canal se cierre y esté vacío
	for idBotella := range ch {
		fmt.Printf("📦 Recibiendo botella [%d]...\n", idBotella)

		// Simula trabajo lento (1 segundo)
		// Mientras esto pasa, el buffer se va llenando porque la fábrica sigue trabajando
		time.Sleep(time.Second * 1)

		fmt.Printf("✅ Botella [%d] EMPACADA\n", idBotella)
		fmt.Println("--------------------------------")
	}
	fmt.Println("--- EMPACADORA TERMINÓ ---")
}

func main() {
	// Creamos el canal con BUFFER de 3
	cintaCh := make(chan int, 3)
	nBotellas := 10

	var wg sync.WaitGroup

	// Solo esperamos a 1 trabajador: La Empacadora.
	// (La fábrica terminará mucho antes, no hace falta esperarla explícitamente)
	wg.Add(1)

	// Lanzamos las goroutines compartiendo el MISMO canal
	go fabricMachine(nBotellas, cintaCh)
	go empacadora(cintaCh, &wg) // Pasamos puntero al WG

	// Esperamos a que la empacadora termine todo
	wg.Wait()
	fmt.Println("Programa finalizado.")
}
