package main

import "fmt"

type Student struct {
	name   string
	group  string
	grades []float64
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
			"4B",
			[]float64{9.3, 9.5, 8.4},
		},
		"02": Student{
			"Matei",
			"4B",
			[]float64{7.7, 6.3, 7.4},
		},
		"03": Student{
			"Toni",
			"4C",
			[]float64{7.4, 9.1, 6.0},
		},
	}

	fmt.Printf("For group %s, average of grades is: %f\n", "4B", generateAvgPerGroup("4B", school))
}
