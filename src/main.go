package main

import (
	"fmt"
	"log"
)

func showParameters(a, b int, name string) {
	log.Println(a, b, name)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func loop() {
}

func isPar(num int) bool {
	return num%2 == 0
}

func main() {

	log.Println("Hola Mundo")

	// Declaracion de constantes

	const pi float64 = 3.14

	log.Println(pi)

	// Declaracion variable Entera

	base := 32
	var altura int = 14
	var area int

	log.Println(base, altura, area)

	// variables no inicializadas

	var isTrue bool
	var numero int
	var decimal float32
	var cadena string

	log.Println(isTrue, numero, cadena, decimal)

	c := 10 + 8i + 2

	log.Println(c)

	nombre := "Nicolas"
	edad := 28

	log.Printf("%s tiene %d años", nombre, edad)

	message := fmt.Sprintf("%s tiene %d años", nombre, edad)

	log.Println(message)

	fmt.Printf("Hello world: %T", nombre)

	value := returnValue(2)
	log.Println("Return value", value)

	value1, value2 := doubleReturn(5)

	log.Println("value1 y value2", value1, value2)

	valor1 := 1
	//valor2 := 2

	if valor1 == 1 {
		log.Println("Es 1")
	} else {
		log.Println("No es 1")
	}

	if isPar(3) {
		log.Println("Es par")
	} else {
		log.Println("Es impar")
	}

	// switch con condicion

	switch modulo := 4 % 2; modulo {

	case 0:
		log.Println("Es par")
	default:
		log.Println("Es impar")
	}

	//switch sin condicion

	valueSwitch := 2

	switch {

	case valueSwitch > 100:
		log.Println("Es mayor a 100")
	case valueSwitch < 0:
		log.Println("Es menor a 0")
	default:
		log.Println("No cumple con ninguna condicion")
	}

	//defer es un keyword que se ejecuta antes que todo termine
	// se usa para cerrar conexiones a db o cuando se abre un archivo

	defer log.Println("hola")
	log.Println("mundo")

	// continue se utilice en caso de que ocurra algun error que tengas controlado
	// y quieres que siga funcionando

	//break detener ejecucion

	for i := 0; i < 10; i++ {

		log.Println(i)
		if i == 2 {
			log.Println("Es dos")
			continue
		}

		if i == 6 {
			log.Println("es 6")
			break
		}

	}

}
