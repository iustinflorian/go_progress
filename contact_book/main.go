// version 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type contact struct {
	name        string
	prefix      string
	phoneNumber string
}

func newContact(name string, prefix string, phoneNumber string) *contact {
	c := contact{name: name, prefix: prefix, phoneNumber: phoneNumber}
	return &c
}

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Println("Local contact book:")

	contactBook := []*contact{}

	for {
		fmt.Print("Add a new contact (type 'no' to cancel):\nName: ")
		name, _ := input.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "no" {
			break
		}

		fmt.Print("Prefix: ")
		prefix, _ := input.ReadString('\n')
		prefix = strings.TrimSpace(prefix)

		fmt.Print("Phone number: ")
		phoneNumber, _ := input.ReadString('\n')
		phoneNumber = strings.TrimSpace(phoneNumber)

		contactBook = append(contactBook, newContact(name, prefix, phoneNumber))
	}

	for _, contact := range contactBook {
		fmt.Println(*contact)
	}

	//contactBook = append(contactBook, newContact("Florian", "+40", "0723234234"))
	// fmt.Println(*contactBook[0])

	// c1 := contact{"Iustin", "+40", "0712123123"}
	// c2 := newContact("Florian", "+40", "0723234234")
	// fmt.Printf("%s -> %s%s\n", c1.name, c1.prefix, c1.phoneNumber)
}
