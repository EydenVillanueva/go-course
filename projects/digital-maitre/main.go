package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
)

type Cliente struct {
	Nombre   string
	Personas int
}

type Mesa struct {
	Numero        int
	Capacidad     int
	Ocupada       bool
	ClienteActual *Cliente
}

const (
	MIN_MESA_CAPACIDAD = 2
	MAX_MESA_CAPACIDAD = 8
	NUM_MESAS          = 10
)

var mapaOcupacion map[int]*Mesa
var listaDeEspera []Cliente
var registroDePagos map[string]int

func init() {
	mapaOcupacion = make(map[int]*Mesa)
	listaDeEspera = make([]Cliente, 0)
	registroDePagos = make(map[string]int)

	for i := 1; i <= NUM_MESAS; i++ {
		capacidad := generateRandomNumber(MIN_MESA_CAPACIDAD, MAX_MESA_CAPACIDAD)
		iMesa := Mesa{
			Numero:        i,
			Capacidad:     capacidad,
			Ocupada:       false,
			ClienteActual: nil,
		}

		mapaOcupacion[i] = &iMesa
	}
}

func generateRandomNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func exitReport() {
	fmt.Println("Generando reporte...")
	counter := 1
	for cliente, total := range registroDePagos {
		fmt.Printf("cuenta: %d cliente: %s Total: %d\n", counter, cliente, total)
		counter++
	}
	fmt.Println("---------------------")
}

func limpiarDisplay() {
	print("\033[H\033[2J")
}

func displayMenu() {
	fmt.Println(`
				--- GESTIÓN DE RESTAURANTE ---
				0. Muestra el menú
				1. Ver estado de mesas y cola
				2. Registrar cliente (Llega gente)
				3. Asignar mesa manual (Saltarse la cola)
				4. Asignar siguiente de la cola a una mesa
				5. Liberar una mesa (Pagar cuenta)
				6. Salir
				> 
			`)
}

func displayStatus() {
	limpiarDisplay()
	fmt.Println("Ocupación...")
	for _, i := range mapaOcupacion {
		var cliente string

		if c := i.ClienteActual; c != nil {
			cliente = i.ClienteActual.Nombre
		} else {
			cliente = "N/A"
		}
		fmt.Printf("numero de mesa: %d capacidad: %d ocupada: %+v cliente: %+v\n",
			i.Numero,
			i.Capacidad,
			i.Ocupada,
			cliente)
	}
	fmt.Println("Lista de espera...")
	for i, cliente := range listaDeEspera {
		fmt.Printf("numero en la cola: %d cliente: %s numero de personas: %d\n",
			i+1,
			cliente.Nombre,
			cliente.Personas)
	}
}

func askForInput[T any](prompt string) (T, error) {
	var v T
	fmt.Print(prompt + ": ")
	_, err := fmt.Scan(&v)
	return v, err
}

func asignarMesaManualmente() {
	fmt.Println()
	idMesa, err := askForInput[int]("Ingrese el número de la mesa")
	if err != nil {
		fmt.Println("Input inválido:", err)
		return
	}

	fmt.Println(" -- Datos del cliente --")

	nombreCliente, err := askForInput[string]("Ingrese nombre del cliente")
	if err != nil {
		fmt.Println("Input inválido:", err)
		return
	}

	npersonas, err := askForInput[int]("Ingrese numero de acompañantes")
	if err != nil {
		fmt.Println("Input inválido:", err)
		return
	}

	cliente := Cliente{
		Nombre:   nombreCliente,
		Personas: npersonas,
	}

	_, err = asignarMesa(idMesa, &cliente)
	if err != nil {
		fmt.Println("Algo falló al asignar la mesa manualmente:", err)
		return
	}

	fmt.Printf("Mesa %d asignada a %s\n", idMesa, nombreCliente)
}

func asignarSigMesa() {
	size := len(listaDeEspera)

	if size == 0 {
		fmt.Println("Lista de espera vacia.. registra nuevos clientes")
		return
	}

	sig := listaDeEspera[0]

	for _, mesa := range mapaOcupacion {
		if !mesa.Ocupada {
			if sig.Personas <= mesa.Capacidad {
				mapaOcupacion[mesa.Numero].Ocupada = true
				mapaOcupacion[mesa.Numero].ClienteActual = &sig
				fmt.Printf("Mesa: %d asignada al cliente: %s", mesa.Numero, sig.Nombre)
				listaDeEspera = listaDeEspera[1:]
				return
			}
		}
	}

	fmt.Printf("No se encontro mesa disponible para cliente: %s", sig.Nombre)
}

func asignarMesa(idMesa int, cliente *Cliente) (bool, error) {
	if cliente == nil {
		return false, errors.New("No se encontro el cliente")
	}

	if idMesa > NUM_MESAS {
		return false, errors.New("No se encontro el numero de mesa")
	}

	mesa, ok := mapaOcupacion[idMesa]

	if !ok {
		return false, errors.New("No se encontro la mesa en el mapaOcupacion")
	}

	if mesa.Ocupada {
		errMessage := fmt.Sprintf("La mesa: %d esta ocupada por el cliente: %+v", mesa.Numero, mesa.ClienteActual)
		return false, errors.New(errMessage)
	}

	if cliente.Personas > mesa.Capacidad {
		errMessage := fmt.Sprintf("Mesa con capacidad insuficiente (%d)", mesa.Capacidad)
		return false, errors.New(errMessage)
	}

	mapaOcupacion[idMesa].Ocupada = true
	mapaOcupacion[idMesa].ClienteActual = cliente
	return true, nil
}

func registrarCliente() (bool, error) {
	fmt.Println()

	fmt.Println(" -- Datos del cliente --")
	nombreCliente, err := askForInput[string]("Ingrese nombre del cliente")
	if err != nil {
		return false, errors.New("Input inválido")
	}

	npersonas, err := askForInput[int]("Ingrese numero de acompañantes")
	if err != nil {
		return false, errors.New("Input inválido")
	}

	isDup := slices.IndexFunc(listaDeEspera, func(c Cliente) bool {
		return c.Nombre == nombreCliente
	})

	if isDup != -1 {
		return false, errors.New("Cliente Duplicado")
	}

	listaDeEspera = append(listaDeEspera, Cliente{
		Nombre:   nombreCliente,
		Personas: npersonas,
	})

	return true, nil
}

func pagarCuenta() (bool, error) {

	idMesa, err := askForInput[int]("Ingrese numero de mesa")

	if err != nil {
		fmt.Println("Numero de mesa invalido")
	}

	// Simulate amount in bill to pay
	cuenta := generateRandomNumber(100, 600)
	mesa, ok := mapaOcupacion[idMesa]

	if !ok {
		errMessage := fmt.Sprintf("Mesa: %s no encontrada", idMesa)
		return false, errors.New(errMessage)
	}

	if !mesa.Ocupada {
		errMessage := fmt.Sprintf("Mesa: %s no esta ocupada", idMesa)
		return false, errors.New(errMessage)
	}

	nombreCliente := mapaOcupacion[idMesa].ClienteActual.Nombre
	registroDePagos[nombreCliente] = cuenta
	mapaOcupacion[idMesa].Ocupada = false
	mapaOcupacion[idMesa].ClienteActual = nil

	return true, nil
}

func startMenu() {
	displayMenu()
loop:
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Print(err)
		}

		switch input {
		case "0":
			displayMenu()
		case "1":
			displayStatus()
			displayMenu()
		case "2":
			_, err := registrarCliente()
			if err != nil {
				fmt.Println("Cliente no pudo ser registrado")
			}
			displayMenu()
		case "3":
			asignarMesaManualmente()
			displayMenu()
		case "4":
			asignarSigMesa()
			displayMenu()
		case "5":
			_, err := pagarCuenta()
			if err != nil {
				fmt.Println("Cuenta no pudo ser pagada")
			}
			displayMenu()
		case "6":
			break loop
		default:
			fmt.Printf("%s no es una opción soportada", input)
			displayMenu()
			continue
		}
	}
}

func main() {
	startMenu()

	defer exitReport()
}
