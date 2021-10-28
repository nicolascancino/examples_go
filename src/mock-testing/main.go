package main

import (
	"fmt"
	"time"
)

type Persona interface {
	GetPerson() string
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Id int
}

type FullTimeEmployee struct {
	Person
	Employee
	TaxRate int
}

type TemporaryEmployee struct {
	*Person
	*Employee
	TaxRate int
}

func (fullTimeEmployee *FullTimeEmployee) GetPerson() string {
	return fullTimeEmployee.Name
}

func (TemporaryEmployee *TemporaryEmployee) GetPerson() string {
	return "Temporary employee"
}

func (person *Person) GetPerson() string {
	return "Normal Person"
}

func (employee *Employee) GetPerson() string {
	return "Employee"
}

func GetPerson(p Persona) {
	fmt.Println(p.GetPerson())
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	e, err := GetEmployeeById(id)

	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e
	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p

	return ftEmployee, nil

}
