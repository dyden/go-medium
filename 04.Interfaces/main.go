package main

import "fmt"

func main() {

	//FULL TIME EMPLOYEE
	ftEmployee := FullTimeEmployee{}
	tEmployee := TemporyEmployee{}
	ftEmployee.Person.Name = "John"
	ftEmployee.Person.Age = 30
	ftEmployee.Employee.Id = 30
	fmt.Println("FULL TIME EMPLOYEE: ", ftEmployee)
	//TEMPORARY EMPLOYEE
	tEmployee.Person.Name = "Jane"
	tEmployee.Person.Age = 25
	tEmployee.Employee.Id = 25
	fmt.Println("TEMPORARY EMPLOYEE: ", tEmployee)
	//GET MESSAGE
	getMessage(ftEmployee)
	getMessage(tEmployee)

}

type PrintInfo interface {
	GetMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
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

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return ftEmployee.Person.Name + " is a full time employee"
}

type TemporyEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee TemporyEmployee) GetMessage() string {
	return ftEmployee.Person.Name + " is a tempory time employee"
}
