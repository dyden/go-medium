package main

import "fmt"

func main() {
	//FULL TIME EMPLOYEE
	fullTimeEmployee := FullTimeEmployee{}
	//INITIALIZE FULL TIME EMPLOYEE
	fullTimeEmployee.Name = "Jhon"
	fullTimeEmployee.Age = 10

	fullTimeEmployee.Id = 1

	fullTimeEmployee.TaxRate = 0.1

	fmt.Println("Name: ", fullTimeEmployee.Name)
	fmt.Println("Age: ", fullTimeEmployee.Age)
	fmt.Println("ID: ", fullTimeEmployee.Id)
	fmt.Println("TaxRate: ", fullTimeEmployee.TaxRate)

}

/***************************************************************
*                          STRUCTS                              *
****************************************************************/
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
	TaxRate float64
}
