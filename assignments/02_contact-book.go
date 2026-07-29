// version 2 - using maps
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type contact struct {
	prefix      string
	phoneNumber string
	isFavourite bool
}

func newContact(prefix string, phoneNumber string) *contact {
	c := contact{prefix: prefix, phoneNumber: phoneNumber, isFavourite: false}
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

	fmt.Print("Do you want to mark it as favourite? (y/n): ")
	choice, _ := input.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "y" {
		contactBook[name].isFavourite = true
		fmt.Printf("Added contact %s as favourite\n\n", name)
		return
	}

	fmt.Printf("Added contact %s\n\n", name)
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

func updateContact(contactBook map[string]*contact, input *bufio.Reader) {
	fmt.Print("Update contact:\nName: ")
	name, _ := input.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("\n")

	_, exists := contactBook[name]
	if !exists {
		fmt.Println("No contact found\n")
		return
	} else {
		fmt.Printf("Contact %s exists. Enter new phone number: ", name)
	}

	newNumber, _ := input.ReadString('\n')
	newNumber = strings.TrimSpace(newNumber)

	for _, info := range contactBook {
		if info.phoneNumber == newNumber {
			fmt.Printf("Contact with phone number %s already exists\n\n", newNumber)
			return
		}
	}

	contactBook[name].phoneNumber = newNumber
	fmt.Printf("Phone number %s saved for contact %s\n\n", newNumber, name)
}

func showContact(contactBook map[string]*contact) {
	if len(contactBook) == 0 {
		fmt.Println("Empty contact book\n")
		return
	}

	fmt.Println("Contact Book:")

	names := make([]string, 0, len(contactBook))
	for name := range contactBook {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		info := contactBook[name]
		fmt.Printf(name + " | " + info.prefix + info.phoneNumber)
		if info.isFavourite {
			fmt.Print(" | Favourite")
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func showMenu(contactBook map[string]*contact, input *bufio.Reader) {
	for {
		fmt.Println("Menu")
		fmt.Println("1. View contacts")
		fmt.Println("2. Add contact")
		fmt.Println("3. Update contact")
		fmt.Println("4. Remove contact")

		fmt.Print("Enter command (type 'no' to cancel): ")

		command, _ := input.ReadString('\n')
		command = strings.TrimSpace(command)
		fmt.Print("\n")

		switch command {
		case "1":
			showContact(contactBook)
		case "2":
			addContact(contactBook, input)
		case "3":
			updateContact(contactBook, input)
		case "4":
			removeContact(contactBook, input)
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
