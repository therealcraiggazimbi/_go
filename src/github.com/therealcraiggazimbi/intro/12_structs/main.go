package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + "and I am " + strconv.Itoa(p.age)
}

// Greeting method (pointer receiver because we are going to change someting)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct

	person1 := Person{firstName: "Deborah", lastName: "Gazimbi", city: "Cape Town", gender: "M", age: 25}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
