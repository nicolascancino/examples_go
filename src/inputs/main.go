package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	operador1, err1 := strconv.Atoi(valores[0])

	// Manejo de error
	if err1 != nil {
		fmt.Println(err1)
	}
	operador2, _ := strconv.Atoi(valores[1])
	fmt.Println(operador1 + operador2)

	// OR

	var operator string
	fmt.Println("Ingrese una operacion")
	fmt.Scanln(&operator)
	fmt.Println(operator)

}
