package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	name   string
	group  string
	grades []float64
}

type ErrStudentNotFound string
type ErrGroupNotFound string

func (e ErrStudentNotFound) Error() string {
	return fmt.Sprintf("student with ID %s not found", string(e))
}
func (e ErrGroupNotFound) Error() string {
	return fmt.Sprintf("group %s not found", string(e))
}

func showStudentInfo(school map[string]*Student, input *bufio.Reader) {
	fmt.Printf("Student ID to search: ")
	ids, _ := input.ReadString('\n')
	ids = strings.TrimSpace(ids)

	student, ok := school[ids]
	if !ok {
		fmt.Println(ErrStudentNotFound(ids))
	}

	sum := 0.0
	for _, grade := range student.grades {
		sum += grade
	}

	sum /= float64(len(student.grades))
	fmt.Printf("Student %s from group %s has an average grade of %f.\n",
		student.name, student.group, sum)
}

/*func genAvgPerGroup(group string, school map[string]*Student) (float64, error) {
	avg := 0.0
	count := 0
	for _, student := range school {
		if student.group == group {
			count++
			sum := 0.0
			for _, grade := range student.grades {
				sum += grade
			}
			avg += sum / float64(len(student.grades))
		}
	}

	if count == 0 {
		return 0, ErrGroupNotFound(group)
	}

	avg /= float64(count)
	return avg, nil
}*/

func showMenu(school map[string]*Student, input *bufio.Reader) {
	for {
		fmt.Println("Menu")
		fmt.Println("1. Add student")
		fmt.Println("2. Update student")
		fmt.Println("3. View student info")
		fmt.Println("4. Generate reports")

		fmt.Print("Enter command (type 'no' to cancel): ")

		command, _ := input.ReadString('\n')
		command = strings.TrimSpace(command)
		fmt.Print("\n")

		switch command {
		case "1":
		case "2":
		case "3":
			showStudentInfo(school, input)
		case "4":
		case "no":
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	school := map[string]*Student{
		"01": &Student{"Andrei", "4b", []float64{9.3, 9.5, 8.4}},
		"02": &Student{"Matei", "4b", []float64{7.7, 6.3, 7.4}},
		"03": &Student{"Toni", "4c", []float64{7.4, 9.1, 6.0}},
	}

	showMenu(school, input)
}
