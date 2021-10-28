package main

import (
	"cursoGo/src/mypackage"
	calc "cursoGo/src/structs/calculator"
	"cursoGo/src/structs/employee"
	"fmt"
)

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	var myCar mypackage.Car
	fmt.Println(myCar)

	var otherCar mypackage.Car

	otherCar.Brand = "Ferrari"
	otherCar.Year = 2020

	fmt.Println(otherCar)

	mypackage.PrintMessage()

	myPc := pc{ram: 16, disk: 200, brand: "MSI"}

	fmt.Println(myPc.String())

	var firstNumber int
	var secondNumber int
	var operator string

	fmt.Println("Ingrese primer numero")
	fmt.Scanln(&firstNumber)

	fmt.Println("Ingrese segundo numero")
	fmt.Scanln(&secondNumber)

	fmt.Println("Ingrese operacion que desea realizar. Ejm: Suma, Resta, Multiplicacion, Division")
	fmt.Scanln(&operator)

	calculator := calc.NewCalculator(firstNumber, secondNumber, operator)

	fmt.Println("Resultado de la operacion: ", calculator.Operation())

	//p := employee.FullTimeEmployee{}
	person := &employee.Person{
		Name: "Nicolas",
		Age:  25,
	}
	empleado := &employee.Employee{
		Id: 1,
	}
	fullTimeEmployee := &employee.FullTimeEmployee{
		Person:   person,
		Employee: empleado,
		TaxRate:  10,
	}
	temporaryEmployee := employee.TemporaryEmployee{}
	employee.GetPerson(empleado)
	employee.GetPerson(person)
	employee.GetPerson(fullTimeEmployee)
	employee.GetPerson(&temporaryEmployee)

}
