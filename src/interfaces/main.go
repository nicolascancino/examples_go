package main

import (
	"fmt"
	"reflect"
)

type animals interface {
	move() string
}

type dog struct {
}

type fish struct {
}

type bird struct {
}

func (dog dog) move() string {
	return "Soy un perro y estoy caminando"
}

func (fish fish) move() string {
	return "Soy un pescado y estoy nadando"
}

func (bird bird) move() string {
	return "Soy un pajaro y estoy volando"
}

type figura2d interface {
	area() float64
	perimetro() float64
}

type cuadrado struct {
	base float64
}

type restangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r restangulo) area() float64 {
	return r.base * r.altura
}

func (c cuadrado) perimetro() float64 {
	return c.base * 4
}
func (r restangulo) perimetro() float64 {
	return r.base * 4
}

func calcular(f figura2d) {
	fmt.Println("Area: ", f.area())

}
func move(animals animals) {
	fmt.Println(animals.move())
}
func main() {

	miCuadrado := cuadrado{base: 4}
	miRestangulo := restangulo{base: 2, altura: 4}

	calcular(miCuadrado)
	calcular(miRestangulo)

	miInterface := []interface{}{"hola", 12, 4.90}
	fmt.Println(miInterface...)

	fmt.Println(reflect.TypeOf(miCuadrado).Name())

	d := dog{}
	f := fish{}
	b := bird{}

	move(d)
	move(f)
	move(b)

}
