package main

import (
	"log"
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id: 1,
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Nicolas",
						Age:  18,
					}, nil
				}

			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "Nicolas",
					Age:  18,
				},
				Employee: Employee{
					Id: 1,
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni
}

func TestGetPerson(t *testing.T) {

	var name = "Nicolas"
	p := &Person{
		Name: "Nicolas",
		Age:  18,
	}
	result := p.GetPerson()
	log.Println(result)
	if result != "Normal Person" {
		t.Errorf("Error, got %s expected %s", p.Name, name)
	}
}
