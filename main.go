package main

import (
	"advance-functions/anonima"
	"advance-functions/operaciones"
	"advance-functions/ordensuperior"
	"advance-functions/recursiva"
	"advance-functions/variadica"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	fmt.Println(variadica.Suma("Tomas", 12, 4, 78, 56))
	fmt.Println(variadica.Suma("Cristina", 10, 20, 30, 40, 50))
	variadica.ImprimirDatos("Hola", 35, "Elias", "Carla", 45.45, "Tomas", -1)

	fmt.Println(recursiva.Factorial(5))

	//funcion anónima
	func() {
		log.Printf("%+v\n", "hola soy una funcion anónima")
	}()

	//la funcion anónima se puede asignar a una variable
	saludo := func(name string) {
		log.Printf("hola %+v, soy una funcion anónima\n", name)
	}

	// y luego usar la variable como si fuera un método, en este caso la función recibe un parámetro
	saludo("Elias")

	//tambien podría crear una función que reciba a su vez un string y una función e internamente la ejecute
	anonima.Saludar("Elias", saludo)

	//es decir estamos pasando las funciones como un valor. Veamos el siguiente ejemplo
	r1 := operaciones.Aplicar(operaciones.Duplicar, 5)
	r2 := operaciones.Aplicar(operaciones.Triplicar, 5)
	fmt.Println(r1, r2)

	//funciones de orden superior: la función Double ejecuta otras funciones
	//Así, en este caso, primero ejecuta Double y luego AddOne | Resultado : 11
	r3 := ordensuperior.Double(operaciones.AddOne, 5)
	log.Printf("Resultado : %+v\n", r3)

}
