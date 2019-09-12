package array

import "fmt"

func Start() {
	//starts [rows][columns]
	a := [2][2]int{ //a 2D array is declared
		{4, 5},
		{7, 85}, // This trailing comma is mandatory
	}
	fmt.Printf("The Array Capacity is = %v\n", cap(a))
	fmt.Printf("The Array Length is = %v\n", len(a))

	// Just like 1D arrays, you don't need to initialize all the elements in a multi-dimensional array.
	// Un-initialized array elements will be assigned the zero value of the array type.
	b := [3][4]float64{
		{1, 3},
		{4.5, -378, 7.744, 2},
		{689, 2, 171},
	}
	fmt.Printf("The Array Capacity is = %v\n", cap(b))
	fmt.Printf("The Array Length is = %v\n", len(b))
}
