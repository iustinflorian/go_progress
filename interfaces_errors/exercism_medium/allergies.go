package main

import "fmt"

type ErrInvalidScore int

func (e ErrInvalidScore) Error() string {
	return fmt.Sprintf("Score value (current: %d) should be greater or equal to 0.", int(e))
}

type Allergen struct {
	Name  string
	Score int
}

var allergenList = []Allergen{
	{"eggs", 1},
	{"peanuts", 2},
	{"shellfish", 4},
	{"strawberries", 8},
	{"tomatoes", 16},
	{"chocolate", 32},
	{"pollen", 64},
	{"cats", 128},
}

func Allergies(score int) ([]string, error) {
	if score < 0 {
		return nil, ErrInvalidScore(score)
	}

	var result []string
	for _, allergen := range allergenList {
		if (score & allergen.Score) != 0 {
			result = append(result, allergen.Name)
		}
	}

	return result, nil
}

func AllergicTo(score int, allergen Allergen) (bool, error) {
	if score < 0 {
		return false, ErrInvalidScore(score)
	}

	if (score & allergen.Score) != 0 {
		return true, nil
	}
	return false, nil
}

func main() {
	sc := 55

	result, err := Allergies(sc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Allergic to %s\n\n", result)

	for _, allergen := range allergenList {
		result, err := AllergicTo(sc, allergen)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(allergen.Name, result)
	}
}
