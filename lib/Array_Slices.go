package lib

import (
	"fmt"
)

func ArrayAndSlices() {
	// Array and Slices#
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("This is my Array... %d\n", numbers)
	//find out the last value
	fmt.Printf("This is the last value in my array is %d\n", numbers[len(numbers)-1])
	// find out the length of array
	fmt.Printf("This is the length of my array %d\n", len(numbers))
	//find the first value in array
	fmt.Printf("This is the first value in  my array %d\n", numbers[0])

	// if we dont know the length of array we can use[...] in the place of [length]like
	//  numberAtInit := [...]int{10,20,30,40,850}

	// Matrix Array
	matrix := [2][3]int{
		{1, 2, 3},
		{5, 6, 7},
		// the [2]representing the the two array vales in the array and [3] is representing the values in array
	}

	fmt.Printf("The MUlti-Dimensional Array is %v\n", matrix)
	//slices are dimensional array
	// allnumbers := numberss[:]
	// firstThree := numbers[0:3]

	//String Array
	fruits := []string{"apple", "banana", "Pineapple"}
	fmt.Printf("This is fruit array %v\n", fruits)
	fruits = append(fruits, "Mango") // if you need to add something else in array
	fmt.Printf("This is fruit array %v\n", fruits)

	//Append In Array(to add something in array)
	moreFruits := []string{"Kiwi", "Strawberry"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("This is my new fruits array %v\n", fruits)

}
