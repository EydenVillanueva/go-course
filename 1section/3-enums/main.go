package main

import "fmt"

// Identificador pre-declarado que actúa como un contador automatico
// para simplificar la creación de constantes numéricas incrementales
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_       = iota // Usamos _ para ignorar el 0
	Enero          // iota es 1
	Febrero        // iota es 2
)

// O también:
const (
	January  = iota + 1 // 0 + 1 = 1
	February            // 1 + 1 = 2
)

type LogLevel int // Aqui creas tu propio tipo llamado LogLevel

const (
	logError LogLevel = iota
	LogWarn
	LogInfo
	logDebug
	logFatal
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

}
