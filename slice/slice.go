package slice

import "fmt"

func Start() {
	primes := [...]int{2, 3, 5, 7, 11, 13} //an array of six elements
	primesSlice := []int{3, 5, 7}          //same array declared as a slice instaed

	var s []int = primes[1:4] // a slice is made from the primes array
	fmt.Println(s)
	fmt.Println(primesSlice)

	//Another example with structs
	b := []struct {
		id   int
		name string
	}{
		{1, "Joe"},
		{2, "John"},
	}

	fmt.Println(b)
	/*
		The variable b is the same as declaring the struct first and then using it as a type for the slice like so
		type Names struct{
			id int
			name string
		}
		b := []Names{
			{1, "Joe"},
			{2, "John"},
		}
	*/
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var a []int
	//The are all equal and equivalent syntax
	a = array[0:]
	fmt.Println(a)
	a = array[:10]
	fmt.Println(a)
	a = array[:]
	fmt.Println(a)
	a = array[0:10]
	fmt.Println(a)
	//
	//Creating a Slice using the make function/keyword
	usingMake := make([]int, 2, 5)
	//where the first params is the data type of the slice
	//2nd param is the length of the slice
	//3rd param is the capacity of the underlining array
	//The make keyword creates a zeroed slice. So if the data type is int it will 0
	fmt.Println(usingMake)
	//Appending Values to a Slice
	//where 1st param is the slice
	//the remaining params will be added to the slice
	appendedSlice := append(usingMake, 22, 78, 95)
	fmt.Println(appendedSlice)

	//Range of a Slice
	//The First Param of Range is the Index
	//The Second Param is a COPY of the value at that index
	//You can escape the index by using a _
	//You can escape the value by omitting it compeletely
	for i, v := range appendedSlice{
		fmt.Printf("The index is %v, The Value is %d \n", i, v)
	}
}
