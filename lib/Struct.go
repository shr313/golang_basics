package lib

import (
	"fmt"
	// "reflect"	//its a built in
)

type Person struct {
	Name   string
	Age    int
	id     int
	Phone  string
	street string
	City   string
}

func AdvanceDS() { //Defines Struct
	pf := fmt.Printf
	pln := fmt.Println

	pln("Advanced Data Structure, Struct, Pointers") //

	person := Person{Name: "Jonathan", Age: 36}
	pf("This is out person %+v\n ", person)

	employee := struct {
		name string
		id   int
	}{
		name: "Alice",
		id:   5321,
	}
	pln("New Employee details are", employee)
	// pln(reflect.TypeOf())
	type Address struct {
		street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}
	contact := Contact{
		Name: "Mark",
		Address: Address{
			street: "123 Main street",
			City:   " Main Town",
		},
	}
	pln("New employee's Detail are", employee)
	pln("New employee's Detail are", contact)
	person.modifyPersonName()
	pln("name before:", person.Name)

	juice := "Mango Juice"
	juicePointer := &juice
	*juicePointer = "Apple Juice"
	fmt.Printf("Value of juice is %v and address is %p\n", juice, juicePointer)
	fmt.Printf("Value of new juice is %v and address is %p\n", juice, juicePointer)
}

// this function is a method attached to Person's Struct
// (person *Person) this method called method receiver
func (person *Person) modifyPersonName() {
	person.Name = "Melkey"
	fmt.Println("inside scope: new name", person.Name)

}
