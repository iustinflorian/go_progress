package main

import (
	"fmt"
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

func showStudentInfo(id string, school map[string]Student) (float64, error) {
	student, ok := school[id]
	if !ok {
		return 0, ErrStudentNotFound(id)
	}

	sum := 0.0
	for _, grade := range student.grades {
		sum += grade
	}

	sum /= float64(len(student.grades))
	return sum, nil
}

func generateAvgPerGroup(group string, school map[string]Student) (float64, error) {
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
}

func main() {
	school := map[string]Student{
		"01": Student{
			"Andrei",
			"4b",
			[]float64{9.3, 9.5, 8.4},
		},
		"02": Student{
			"Matei",
			"4b",
			[]float64{7.7, 6.3, 7.4},
		},
		"03": Student{
			"Toni",
			"4c",
			[]float64{7.4, 9.1, 6.0},
		},
	}

	studentId := "03"
	studentAvg, err := showStudentInfo(studentId, school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Student %s from group %s has an average grade of %f.\n", school[studentId].name, school[studentId].group, studentAvg)
	}

	group := "4B"
	group = strings.ToLower(group)
	avg, err := generateAvgPerGroup(group, school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("For group %s, average of grades is: %f\n", group, avg)
	}

}
