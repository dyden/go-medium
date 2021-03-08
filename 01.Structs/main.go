package main

import "fmt"

func main() {
	employe := Employee{}
	//DEFAULT VALUE OF EMPLOYEE {0, "", 0}
	fmt.Println(employe)
	//SET ID OF EMPLOYEE
	employe.SetId(1)
	//SET NAME OF EMPLOYEE
	employe.SetName("Jhon")
	//SET AGE OF EMPLOYEE
	employe.SetAge(25)
	//VALUE OF EMPLOYEE AFTER SETTING {1, "Jhon", 25}
	fmt.Println(employe)
	//PRINT ATRIBUTES OF EMPLOYEE
	fmt.Println(employe.GetId(), employe.GetName(), employe.GetAge())

}

/***************************************************************
*                          STRUCTS                              *
****************************************************************/

type Employee struct {
	id   int
	name string
	age  int
}

//GET ID
func (e *Employee) GetId() int {
	return e.id
}

//GET NAME
func (e *Employee) GetName() string {
	return e.name
}

//GET AGE
func (e *Employee) GetAge() int {
	return e.age
}

//SET ID
func (e *Employee) SetId(id int) {
	e.id = id
}

//SET NAME
func (e *Employee) SetName(name string) {
	e.name = name
}

//SET AGE
func (e *Employee) SetAge(age int) {
	e.age = age
}
