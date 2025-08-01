package main

import "fmt"

type inputError struct {
	message string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func main() {
	err := validate("","")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err, ok)
			fmt.Printf("missing field is %s\n", err.getMissingField())
			}
	}
}


func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "name is mandatory", missingField: "name"}
	}
	if gender == "" {
		return &inputError{message: "gender is mandatory", missingField: "gender"}
	}
	return nil
}