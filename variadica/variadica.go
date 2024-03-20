package variadica

import (
	"fmt"
	"log"
)

// Función Variádica, es decir que recibe "n" parametros, eso lo logro con los "..."
func Suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)
	return total
}

// imprimir datos, al usar interface indico que puede ser n parametros de distintos tipos
func ImprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		log.Printf("%T - %+v\n", dato, dato)
	}
}
