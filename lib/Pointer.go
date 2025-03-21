package lib

import "fmt"

func Pointer() {
	Alice := Person{
		Name:   "Alice",
		Age:    36,
		id:     5321,
		Phone:  "12345678",
		street: "123 Main street",
		City:   " Main Town",
	}
	fmt.Printf("Name before: %v\n", Alice.Name)
	ChangeName(&Alice, "Johnathan ") // (&) is being used as address of the variable
	fmt.Printf("Name After: %v\n", Alice.Name)
}

func ChangeName(person *Person, newName string) {
	person.Name = newName
	fmt.Printf("Name inside Function Scope: %v\n", person.Name)
}
