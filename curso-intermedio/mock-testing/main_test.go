package main

import "testing"

func TestGetFulltimeEmployeeByID(t *testing.T) {
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
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDni = func(dni string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "John Doe",
					Age:  35,
					DNI:  "1",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonById := GetPersonByDni

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeByID(test.id, test.dni)
		if err != nil {
			t.Errorf("Error al obtener empleado")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDni = originalGetPersonById
}
