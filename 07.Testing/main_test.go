package main

import (
	"testing"
)

/***************************************************************
*                        TEST FOR SUM()                        *
****************************************************************/
//YOU CAN USE 'go test -v' TO SEE THE OUTPUT OF THE TEST CASES

func TestSum(t *testing.T) {

	total := sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}
	//TEST CASES
	cases := []struct{ a, b, n int }{{1, 2, 3}, {2, 3, 5}, {3, 4, 8}}
	for _, test := range cases {
		result := sum(test.a, test.b)
		if result != test.n {
			t.Errorf("Sum was incorrect, got %d expected %d", result, test.n)
		}
	}
}

/***************************************************************
*                   TESTING USING MOCKS                        *
****************************************************************/
func TestGetFullTimeEmployeeByID(t *testing.T) {
	cases := []struct {
		ID               int
		DNI              string
		mockFunc         func()
		expectedEMployee FullTimeEmployee
	}{
		{
			ID:  1,
			DNI: "1234",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{ID: 1, Position: "CEO"}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{Name: "Jhon", Age: 25, DNI: "1234"}, nil
				}
			},
			expectedEMployee: FullTimeEmployee{
				Person:   Person{Name: "Jhon", Age: 20, DNI: "1234"},
				Employee: Employee{ID: 1, Position: "CEO"},
			},
		},
	}
	originalGetEmployeeByID := GetEmployeeByID
	originalGetPersonByDNI := GetPersonByDNI
	for _, test := range cases {
		test.mockFunc()
		ftEmployee, err := GetFullTimeEmployeeByID(test.ID, test.DNI)
		if err != nil {
			t.Errorf("Error when getting Employee: %s", err)
		}
		if ftEmployee != test.expectedEMployee {
			t.Errorf("Employee was incorrect, got %v expected %v", ftEmployee, test.expectedEMployee)
		}
	}
	GetEmployeeByID = originalGetEmployeeByID
	GetPersonByDNI = originalGetPersonByDNI
}

/***************************************************************
*                        TEST COVERAGE                         *
****************************************************************/

//YOU CAN USE 'go test -cover' FOR RUN COVERAGE or 'go test -coverprofile=cover.out' FOR SAVE COVERAGE
//USE 'go tool cover -html=cover.out' FOR VIEW COVERAGE ON YOUR WEB BROWSER

/***************************************************************
*                           PROFILING                           *
****************************************************************/

//$ sudo apt install graphviz -> Linux
//$ brew install graphviz -> MacOS

//YOU CAN USE 'go test -cpuprofile=cpu.out' FOR RUN PROFILING OR 'go test -cpuprofile=cpu.out -bench=BenchmarkSum' FOR RUN BENCHMARK
//USE 'go tool pprof cpu.out' FOR VIEW PROFILING AND USE '(pprof) web'TO VIEW PROFILING ON YOUR WEB BROWSER OR '(pprof) pdf' FOR SAVE PROFILING ON PDF
