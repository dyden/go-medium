package main

import "fmt"

func main() {
	//INSTANCES OF EMPLOYE

	//CRETE EMPLOYEE WITH DEFAULT VALUES
	employee1 := Employee{}
	fmt.Println(employee1)
	//CREATE INSTANCE OF EMPLOYEE INITIALIZING WITH VALUES
	employee2 := Employee{ID: 1, Name: "Jhon", Age: 25}
	fmt.Println(employee2)
	//CREATE A POINTER TO EMPLOYEE
	employee3 := new(Employee)
	fmt.Println(*employee3)
	//ANOTHER WAY
	employee4 := NewEmployee(1, "Jhon", 25)
	fmt.Println(*employee4)

}

/***************************************************************
*                          STRUCTS                              *
****************************************************************/

type Employee struct {
	ID   int
	Name string
	Age  int
}

func NewEmployee(ID int, name string, age int) *Employee {
	return &Employee{
		ID:   ID,
		Name: name,
		Age:  age,
	}
}

//GET ID
func (e *Employee) GetID() int {
	return e.ID
}

//GET NAME
func (e *Employee) GetName() string {
	return e.Name
}

//GET AGE
func (e *Employee) GetAge() int {
	return e.Age
}

//SET ID
func (e *Employee) SetID(id int) {
	e.ID = id
}

//SET NAME
func (e *Employee) SetName(name string) {
	e.Name = name
}

//SET AGE
func (e *Employee) SetAge(age int) {
	e.Age = age
}
