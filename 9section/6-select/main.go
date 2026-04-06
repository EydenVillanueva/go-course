package main

import (
	"fmt"
	"time"
)

func main() {

	canalTecnico := make(chan string)
	canalPeligro := make(chan string)

	go func(canalTecnico chan<- string) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			canalTecnico <- "Navegación estable"
		}
	}(canalTecnico)

	go func(canalPeligro chan<- string) {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 2)
			canalPeligro <- "ASTEROIDE DETECTADO!"
		}
	}(canalPeligro)

	for {
		select {
		case msgTecnico := <-canalTecnico:
			fmt.Println(msgTecnico)
		case msgPeligro := <-canalPeligro:
			fmt.Println(msgPeligro)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout: Nadie ha dicho nada en 5 segundos")
			return
		}
	}

}
