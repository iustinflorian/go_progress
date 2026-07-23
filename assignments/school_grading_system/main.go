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

type School map[string]*Student

type ErrStudentNotFound string
type ErrGroupNotFound string

func (e ErrStudentNotFound) Error() string {
	return fmt.Sprintf("student with ID %s not found", string(e))
}
func (e ErrGroupNotFound) Error() string {
	return fmt.Sprintf("group %s not found", string(e))
}

func (s Student) avgGrade() float64 {
	sum := 0.0
	for _, grade := range s.grades {
		sum += grade
	}
	return sum / float64(len(s.grades))
}

/*func (school School) generateID() string {
	var newID, lastID int
	lastID = 0

	for ids := range school {
		var idsNum int
		fmt.Sscanf(ids, "%d", &idsNum)
		if idsNum > lastID {
			lastID = idsNum
		}
	}

	newID = lastID + 1
	return fmt.Sprintf("%02d", newID)
}*/

func (school School) showStudentInfo(input *bufio.Reader) {
	fmt.Printf("Student ID to search: ")
	ids, _ := input.ReadString('\n')
	ids = strings.TrimSpace(ids)

	student, ok := school[ids]
	if !ok {
		fmt.Printf("Error %s\n", ErrStudentNotFound(ids))
		return
	}

	fmt.Printf("Student %s from group %s has an average grade of %.2f.\n",
		student.name, student.group, student.avgGrade())
}

func (school School) genAvgPerGroup(input *bufio.Reader) {
	fmt.Printf("Group ID: ")
	idg, _ := input.ReadString('\n')
	idg = strings.TrimSpace(idg)

	avg := 0.0
	count := 0
	for _, student := range school {
		if student.group == idg {
			count++
			avg += student.avgGrade()
		}
	}
	if count == 0 {
		fmt.Printf("Error %s\n", ErrGroupNotFound(idg))
		return
	}
	avg /= float64(count)

	fmt.Printf("The average grade for group %s is %.2f\n", idg, avg)
}

func (school School) genStudentOrderByAvg() {
	// todo: check if school is empty

	var ids []string
	for id := range school {
		ids = append(ids, id)
	}

	sort.Slice(ids, func(i, j int) bool {
		return school[ids[i]].avgGrade() > school[ids[j]].avgGrade()
	})

	fmt.Printf(" Nr.| Grade | Name\n--------------------\n")
	for i, id := range ids {
		student := school[id]
		fmt.Printf("%3d | %5.2f | %s\n", i+1, student.avgGrade(), student.name)
	}
}

/*func (school School) addStudent(input *bufio.Reader) {
	fmt.Printf("\n----[ Add a new student ]----")
	id := school.generateID()
}*/

func (school School) showMenu(input *bufio.Reader) {
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
			school.showStudentInfo(input)
		case "2":
			school.reportsMenu(input)
		// case "3":
		// case "4":
		case "no":
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}

func (school School) reportsMenu(input *bufio.Reader) {
	fmt.Println("\n---------[ Generate a report ]--------")
	fmt.Println("1. Average grade for a particular group")
	fmt.Println("2. Students list ordered by average grade")
	fmt.Print("Enter command (type 'no' for main menu): ")

	command, _ := input.ReadString('\n')
	command = strings.TrimSpace(command)

	switch command {
	case "1":
		school.genAvgPerGroup(input)
	case "2":
		school.genStudentOrderByAvg()
	case "no":
		return
	default:
		fmt.Println("Invalid command.")
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	s := School{
		"01": &Student{"Andrei", "4b", []float64{9.3, 9.5, 8.4}},
		"02": &Student{"Matei", "4b", []float64{7.7, 6.3, 7.4}},
		"03": &Student{"Toni", "4c", []float64{7.4, 9.1, 6.0}},
	}

	s.showMenu(input)
}
