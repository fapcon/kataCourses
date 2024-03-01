package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func main() {
 fmt.Println(GenerateFakeData())
}

func GenerateFakeData() string {
	var name, address, phone, email string
	name = fake.FullName()
	address = fake.StreetAddress()
	phone = fake.Phone()
	email = fake.EmailAddress()
	result := "Name: " +name + "\n" + "Address: " + address + "\n" + "Phone: " + phone + "\n" + "Email: " + email
	return result
}
