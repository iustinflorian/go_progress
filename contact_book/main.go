// version 2 - using maps
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type contact struct {
	prefix      string
	phoneNumber string
}

func newContact(prefix string, phoneNumber string) *contact {
	c := contact{prefix: prefix, phoneNumber: phoneNumber}
	return &c
}

func addContact(contactBook map[string]*contact, input *bufio.Reader) {
	fmt.Print("Add contact:\nName: ")
	name, _ := input.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Prefix: ")
	prefix, _ := input.ReadString('\n')
	prefix = strings.TrimSpace(prefix)

	fmt.Print("Phone number: ")
	phoneNumber, _ := input.ReadString('\n')
	phoneNumber = strings.TrimSpace(phoneNumber)

	_, exist := contactBook[name]
	if exist {
		fmt.Printf("Contact %s already exists\n\n", name)
		return
	}

	for _, info := range contactBook {
		if info.phoneNumber == phoneNumber {
			fmt.Printf("Contact with phone number %s already exists\n\n", phoneNumber)
			return
		}
	}

	contactBook[name] = newContact(prefix, phoneNumber)
	fmt.Print("\n")
}

func removeContact(contactBook map[string]*contact, input *bufio.Reader) {
	fmt.Print("Remove contact:\nName: ")
	name, _ := input.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("\n")

	_, exists := contactBook[name]
	if !exists {
		fmt.Println("No contact found")
		return
	}

	delete(contactBook, name)
	fmt.Printf("Contact %s removed\n\n", name)
}

func showContact(contactBook map[string]*contact) {
	fmt.Println("Contact Book:")
	for name, info := range contactBook {
		fmt.Printf(name + " -> " + info.prefix + info.phoneNumber)
	}
	fmt.Print("\n\n")
}

func showMenu(contactBook map[string]*contact, input *bufio.Reader) {
	for {
		fmt.Println("Menu")
		fmt.Println("1. Add contact")
		fmt.Println("2. Remove contact")
		fmt.Println("3. View contacts")
		fmt.Print("Enter command (type 'no' to cancel): ")

		command, _ := input.ReadString('\n')
		command = strings.TrimSpace(command)
		fmt.Print("\n")

		switch command {
		case "1":
			addContact(contactBook, input)
		case "2":
			removeContact(contactBook, input)
		case "3":
			showContact(contactBook)
		case "no":
			return
		default:
			fmt.Println("Invalid command\n")
		}
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	contactBook := make(map[string]*contact)

	showMenu(contactBook, input)
}
