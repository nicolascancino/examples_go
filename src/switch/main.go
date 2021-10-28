package main

import (
	"fmt"
	"strings"
)

func main() {
	var operation string
	var firstValue int
	var secondValue int
	var result int
	fmt.Println("Ingrese una operacion, Ej: Suma, Resta, Multiplicacion, Division")
	fmt.Scanln(&operation)
	fmt.Println("Ingrese el primer valor")
	fmt.Scanln(&firstValue)
	fmt.Println("Ingrese el segundo valor")
	fmt.Scanln(&secondValue)

	switch strings.ToLower(operation) {
	case "suma":
		result = firstValue + secondValue
	case "resta":
		result = firstValue - secondValue
	case "multiplicacion":
		result = firstValue * secondValue
	case "division":
		result = firstValue / secondValue
	default:
		fmt.Println("Operador no aceptado")
	}

	fmt.Printf("Resultado: %d\n", result)
}
