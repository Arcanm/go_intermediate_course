package main

import "testing"

func TestGetFultimeEmployeeById(t *testing.T) {
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

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "1",
						Name: "Andres",
						Age:  25,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person:   Person{DNI: "1", Name: "Andres", Age: 25},
				Employee: Employee{Id: 1, Position: "CEO"},
			},
		},
	}

	// Save the original functions
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error GOT %d expected %d", ft.Age, test.expectedEmployee.Age)
		}

		// Restore the original functions
		GetEmployeeById = originalGetEmployeeById
		GetPersonByDNI = originalGetPersonByDNI
	}
}
