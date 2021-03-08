package main

import "time"

func sum(x, y int) int {
	return x + y
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonaaci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonaaci(n-1) + Fibonaaci(n-2)
}

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	ID       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Persona Where ...
	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeByID(id)
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
