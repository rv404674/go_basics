package struct_basics 

import "fmt"

/*
Date - 21 March, 2020
Coursera - Structs
*/

func main(){

	// Struct is a composite date type, much like a class in C++
	type Person struct{
		name string
		address string
		phone_number string
}

	var p1 Person

	// to set the values
	p1.name = "Joe"
	x = p1.address

	p2 := new(Person) //initialized all fields to zero

	//another way to initialize
	p3 := Person(name: "Rahul", address: " a st", phone_number: "1234")

}