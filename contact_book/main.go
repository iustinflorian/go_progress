// version 1
package main

import "fmt"

type contact struct {
	name        string
	phoneNumber string
}

func newContact(name string, phoneNumber string) *contact {
	c := contact{name: name, phoneNumber: phoneNumber}
	return &c
}

func main() {

	fmt.Println("Local contact book.")

	c1 := contact{"Iustin", "0712123123"}

	fmt.Println(c1.phoneNumber)

}
