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

func addContact(contactBook []*contact, input *bufio.Reader) []*contact {
	fmt.Print("Add contact:\nName: ")
	name, _ := input.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Prefix: ")
	prefix, _ := input.ReadString('\n')
	prefix = strings.TrimSpace(prefix)

	fmt.Print("Phone number: ")
	phoneNumber, _ := input.ReadString('\n')
	phoneNumber = strings.TrimSpace(phoneNumber)

	fmt.Print("\n")
	contactBook = append(contactBook, newContact(name, prefix, phoneNumber))

	return contactBook
}

func showContact(contactBook []*contact) {
	fmt.Println("Contact Book:")
	for i, contact := range contactBook {
		fmt.Printf("%d. %s -> %s%s\n", i+1, contact.name, contact.prefix, contact.phoneNumber)
	}
	fmt.Print("\n")
}

func showMenu(contactBook []*contact, input *bufio.Reader) {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add contact:")
		fmt.Println("2. View contacts:")
		fmt.Print("Enter command (type 'no' to cancel): ")

		command, _ := input.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "1":
			contactBook = addContact(contactBook, input)
		case "2":
			showContact(contactBook)
		default:
			return
		}

		fmt.Print("\n")
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	contactBook := []*contact{}

	showMenu(contactBook, input)
}
