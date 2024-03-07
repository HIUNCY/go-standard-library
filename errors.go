package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Validation error!")
	notFoundError   = errors.New("Not found error!")
)

func getById(id string) error {
	if id == "" {
		return validationError
	} else if id != "Ajay" {
		return notFoundError
	} else {
		return nil
	}
}

func main() {
	err := getById("")
	if errors.Is(err, validationError) {
		fmt.Println("Validation error!")
	} else if errors.Is(err, notFoundError) {
		fmt.Println("Not found error!")
	} else {
		fmt.Println("No error detected!") // success
	}
}
