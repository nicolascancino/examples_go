package employee

import "fmt"

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
	*Person
	*Employee
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
