package main

import "fmt"

// Define a contactinfo struct
type ContactInfo struct {
	email   string
	city    string
	zipcode int
}

// Define a person struct
type Person struct {
	firstName   string
	lastName    string
	contactinfo ContactInfo
	// ContactInfo -- This is also a valid syntax
}

func (p Person) print() {
	fmt.Printf("\n----------------\n")
	fmt.Printf("%v", p)
	fmt.Printf("\n----------------\n")
}

func (p Person) updateName(firstName string) Person {
	p.firstName = firstName

	return p
}

type location struct {
	longitude float64
	latitude  float64
}

func main() {
	newYork := location{
		latitude:  40.73,
		longitude: -73.93,
	}

	newYork.changeLatitude()

	fmt.Println(newYork)
	name := "Bill"

	fmt.Println(*&name)
}

func (lo *location) changeLatitude() {
	(*lo).latitude = 41.0
}

// func main() {
// 	// Create a struct by passing arguments in order
// 	amitesh := Person{
// 		"Amitesh",
// 		"Rai",
// 		ContactInfo{
// 			email:   "amiteshrai@outlook.com",
// 			city:    "Banglore",
// 			zipcode: 562107,
// 		},
// 	}
// 	amitesh.print()

// 	// Create a struct by passing arguments as key-value pairs
// 	amitesh = Person{
// 		lastName:  "Rai",
// 		firstName: "Amitesh",
// 		contactinfo: ContactInfo{
// 			email:   "amiteshrai@outlook.com",
// 			city:    "Banglore",
// 			zipcode: 562107,
// 		},
// 	}
// 	fmt.Print(amitesh.contactinfo)

// 	// Create a struct by passing arguments and omitting few
// 	amitesh = Person{
// 		lastName:  "Rai",
// 		firstName: "Amitesh",
// 	}
// 	amitesh.print()
// 	fmt.Println(amitesh.contactinfo) // Assign "" as default, called Zero Value in Go

// 	// Create a struct
// 	var sinchan Person
// 	fmt.Println(sinchan)
// 	fmt.Printf("%+v", sinchan)
// 	amitesh.updateName(" Amitesh ")
// 	amitesh.print()
// }
