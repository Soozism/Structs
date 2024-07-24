package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "jim",
		lastName: "party",
		contactInfo: contactInfo{
			email: "ali@j.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v",jim)
}