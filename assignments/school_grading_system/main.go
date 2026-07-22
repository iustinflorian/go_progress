package main

import (
	"fmt"
	"strings"
)

// todo: functions that show info to use ID instead of passing the struct

type Student struct {
	name   string
	group  string
	grades []float64
}

func showStudentInfo(student Student) {
	sum := 0.0

	for _, grade := range student.grades {
		sum += grade
	}

	sum /= float64(len(student.grades))
	fmt.Printf("Student %s from group %s has an average grade of %f.\n", student.name, student.group, sum)
}

func generateAvgPerGroup(group string, school map[string]Student) float64 {
	avg := 0.0
	count := 0.0

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

	avg /= count
	return avg
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

	showStudentInfo(school["01"])

	group := "4B"
	group = strings.ToLower(group)
	fmt.Printf("For group %s, average of grades is: %f\n", group, generateAvgPerGroup("4b", school))
}
