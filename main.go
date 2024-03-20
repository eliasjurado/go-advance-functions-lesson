package main

import (
	"advance-functions/variadica"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	fmt.Println(variadica.Suma("Tomas", 12, 4, 78, 56))
	fmt.Println(variadica.Suma("Cristina", 10, 20, 30, 40, 50))
	variadica.ImprimirDatos("Hola",35,"Elias","Carla",45.45,"Tomas",-1)
}

