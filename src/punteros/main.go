package main

import (
	pcDomain "cursoGo/src/punteros/pc"
	"fmt"
)

func main() {

	a := 50
	b := &a

	// & obtener direccion de memoria
	// * obtener valor de la direccion de memoria

	fmt.Println(a)
	fmt.Println(&b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pcDomain.New(16, 200, "msi")
	fmt.Println(myPc)
	myPc.Ping()
	myPc.DuplicateRam()
	fmt.Println(myPc)
	myPc.SetRam(64)
	fmt.Println(myPc)

	y := 25
	cambiarValor(&y)
	fmt.Println(y)

}

func cambiarValor(a *int) {
	*a = 30
}
