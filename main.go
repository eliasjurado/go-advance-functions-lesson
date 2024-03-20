package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	fmt.Println(suma("Tomas", 12, 4, 78, 56))
	fmt.Println(suma("Cristina", 10, 20, 30, 40, 50))
	imprimirDatos("Hola",35,"Elias","Carla",45.45,"Tomas",-1)
}

// Función Variádica, es decir que recibe "n" parametros, eso lo logro con los "..."
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)
	return total
}

// imprimir datos, al usar interface indico que puede ser n parametros de distintos tipos
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		log.Printf("%T - %+v\n", dato, dato)
	}
}
