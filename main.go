package main

//  import format package
import (
	"fmt"
)

// Main Function
func main(){

	// control Structure
	age := 30
	if age >= 18 {
		fmt.Println(" You are an adult,your ticket will be £30  ")
	}else if age>= 13{
		fmt.Println("You are an teenager,your ticket will be £18")
	}else {
		fmt.Println("You are a child,your ticket will be £12")
	}
	// Switch Statement
day := "Tuesday"
switch day{
	case "Monday":
	fmt.Printf(" its %s Start of the week.\n ",day)
	case "Tuesday","Wednesday","Thursday":
	fmt.Printf("It's %s midweek.\n ",day)
	case "Friday":
		fmt.Printf("It's a %s funday\n ",day)
		default:
		fmt.Printf("it's the %s weekend\n",day)
}
// For Loop
for i := 0; i <5; i++ {
	fmt.Printf("This is... %d\n", i)
}
// While Loop
 counter := 0
 for counter <5{
	fmt.Printf("Count....%d\n", counter)
	counter++ // use ++(condition) with variable otherwise it'll never stop.
 }

 //infinite Loop
iteration := 0
 for{
	if iteration >3 {
		fmt.Printf("Iteration is %d\n",iteration)
		break//some condition is met
	}
	iteration++
 }

 // Array and Slices#
 numbers := [5]int{10,20,30,40,50}
 fmt.Printf("This is my Array... %d\n", numbers)
 //find out the last value
 fmt.Printf("This is the last value in my array is %d\n",numbers[len(numbers)-1])
 // find out the length of array
 fmt.Printf("This is the length of my array %d\n", len(numbers))
 //find the first value in array
 fmt.Printf("This is the first value in  my array %d\n", numbers[0])

 // if we dont know the length of array we can use[...] in the place of [length]like
//  numberAtInit := [...]int{10,20,30,40,850}

// Matrix Array
matrix := [2][3]int{
	{1,2,3},
	{5,6,7},
// the [2]representing the the two array vales in the array and [3] is representing the values in array
}

fmt.Printf("The MUlti-Dimensional Array is %v\n",matrix)
//slices are dimensional array
// allnumbers := numberss[:]
// firstThree := numbers[0:3]

//String Array
fruits :=[]string{"apple","banana","Pineapple"}
fmt.Printf("This is fruit array %v\n", fruits)
fruits = append(fruits, "Mango")// if you need to add something else in array
fmt.Printf("This is fruit array %v\n", fruits)

moreFruits := []string{"Kiwi", "Strawberry"}
fruits = append(fruits, moreFruits...)
fmt.Printf("This is my new fruits array %v\n", fruits)

// make statements / pre allocated memory of slices

scores := make([]int,3) //slice with the length 3 and capacity 3
scoresWithExplicitCapicity := make([]int, 3, 5)//slice with the length 3 and capacity 5
fmt.Println("scoreWithExplicitCapacity",len(scoresWithExplicitCapicity),cap(scoresWithExplicitCapicity))
fmt.Println("scoreWithExplicitCapacity",len(scores),cap(scores))
// length is the current amount of elements present in the array or slice.
//capacity is the total number of element that can have/ can be added to the array.

for index, value := range numbers{
	fmt.Printf("index %d and value %d\n",index, value)
}
// we use(-,) to void any variable
// or _, value := range numbers{
// 	fmt.Printf("index %d and value %d\n", value)


capitalCities := map[string]string{
	"USA" : "Washington D.C.",
	"Italy" : "Rome",
	"England" : "London",
}
fmt.Println(capitalCities["Italy"])
//if we want to anything exits
capital, exits := capitalCities["Germany"]
if exits { // exits keyword only use in the map
	fmt.Printf("This is the capital of %v\n",capital)
} else{
		fmt.Println("This does not exits.")
	}
}

// add function
func add(a int, b int) int {
	return a+b
}

// Multiple Use of Variable in function
func calculateSumAndProduct(a,b int)(int, int){
	return a+b, a*b
}
