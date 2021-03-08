package main

import (
	"fmt"
)

func main() {

	//ANONYMOUS FUNCTION
	x := 10
	y := func() int {
		return x * 2
	}
	fmt.Println(x, y())
	//FUNCTION WITH UNKNOWN PARAMETERS NUMBER
	sum1 := sum(1, 2, 3)
	sum2 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("SUM 1: ", sum1)
	fmt.Println("SUM 2: ", sum2)
	//FUNCTION WITH RETURN VALUES
	double, tripe, quad := getValue(10)
	fmt.Println("DOUBLE: ", double)
	fmt.Println("TRIPE: ", tripe)
	fmt.Println("QUAD: ", quad)

}

/***************************************************************
*                      FUNCTIONS                               *
****************************************************************/
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
func getValue(value int) (double int, triple int, quad int) {
	double = value * 2
	triple = value * 3
	quad = value * 4
	return
}
