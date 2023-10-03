package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	FirstName string
	LastName  string
	Contact   Contact
}

func (employee Employee) printInfo() {
	fmt.Println("Demo Function Fot Struct")
	fmt.Println("Employee: ", employee.FirstName, employee.LastName)
	fmt.Println("Email: ", employee.Contact.Email)
	fmt.Println("Phone: ", employee.Contact.Phone)
}
func main() {
	//person1 := Person{FirstName: "John", LastName: "Doe", Age: 30}
	//
	//fmt.Println("First Name: ", person1.FirstName)
	//fmt.Println("Last Name: ", person1.LastName)
	//fmt.Println("Age: ", person1.Age)
	//fmt.Println()
	//
	//var person2 Person
	//person2 = Person{FirstName: "Rose", LastName: "Vim", Age: 28}
	//fmt.Println("First Name: ", person2.FirstName)
	//fmt.Println("Last Name: ", person2.LastName)
	//fmt.Println("Age: ", person2.Age)

	employee := Employee{
		FirstName: "Lee",
		LastName:  "Nguyen",
		Contact: Contact{
			Email: "lee.nguyen@gmail.com",
			Phone: "234494838483",
		},
	}

	//fmt.Println("Employee: ", employee.FirstName, employee.LastName)
	//fmt.Println("Email: ", employee.Contact.Email)
	//fmt.Println("Phone: ", employee.Contact.Phone)
	employee.printInfo()

}
