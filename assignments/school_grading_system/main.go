package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func avgGrade(grades []float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func showStudentInfo(school map[string]*Student, input *bufio.Reader) {
	fmt.Printf("Student ID to search: ")
	ids, _ := input.ReadString('\n')
	ids = strings.TrimSpace(ids)

	student, ok := school[ids]
	if !ok {
		fmt.Printf("Error %s\n", ErrStudentNotFound(ids))
		return
	}

	fmt.Printf("Student %s from group %s has an average grade of %.2f.\n",
		student.name, student.group, avgGrade(student.grades))
}

func genAvgPerGroup(school map[string]*Student, input *bufio.Reader) {
	fmt.Printf("Group ID: ")
	idg, _ := input.ReadString('\n')
	idg = strings.TrimSpace(idg)

	avg := 0.0
	count := 0
	for _, student := range school {
		if student.group == idg {
			count++
			avg += avgGrade(student.grades)
		}
	}
	if count == 0 {
		fmt.Printf("Error %s\n", ErrGroupNotFound(idg))
		return
	}
	avg /= float64(count)

	fmt.Printf("The average grade for group %s is %.2f\n", idg, avg)
}

func genStudentOrderByAvg(school map[string]*Student) {
	// todo: check if school is empty

	var ids []string
	for id := range school {
		ids = append(ids, id)
	}

	sort.Slice(ids, func(i, j int) bool {
		return avgGrade(school[ids[i]].grades) > avgGrade(school[ids[j]].grades)
	})

	fmt.Printf(" Nr.| Grade | Name\n--------------------\n")
	for i, id := range ids {
		student := school[id]
		fmt.Printf("%3d | %5.2f | %s\n", i+1, avgGrade(student.grades), student.name)
	}
}

func showMenu(school map[string]*Student, input *bufio.Reader) {
	for {
		fmt.Println("\n----[ Main menu ]----")
		fmt.Println("1. View student info")
		fmt.Println("2. Generate reports")
		// fmt.Println("3. Add student")
		// fmt.Println("4. Update student")
		fmt.Print("Enter command (type 'no' to cancel): ")

		command, _ := input.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "1":
			showStudentInfo(school, input)
		case "2":
			reportsMenu(school, input)
		// case "3":
		// case "4":
		case "no":
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}

func reportsMenu(school map[string]*Student, input *bufio.Reader) {
	fmt.Println("\n---------[ Generate a report ]--------")
	fmt.Println("1. Average grade for a particular group")
	fmt.Println("2. Students list ordered by average grade")
	fmt.Print("Enter command (type 'no' for main menu): ")

	command, _ := input.ReadString('\n')
	command = strings.TrimSpace(command)

	switch command {
	case "1":
		genAvgPerGroup(school, input)
	case "2":
		genStudentOrderByAvg(school)
	case "no":
		showMenu(school, input)
	default:
		fmt.Println("Invalid command.")
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
