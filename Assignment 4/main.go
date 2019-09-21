package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Name   string
	Gender string
	Age    int
}

type error struct {
	message string
}

func testValidate(t *testing.T) {
	noError := error{}

	p := Person{
		Name:   "",
		Gender: "Male",
		Age:    24,
	}
	test := p.Validate()
	if test == noError {
		t.Errorf("Validation should give empty name error message")
	}

	p = Person{
		Name:   "John",
		Gender: "",
		Age:    24,
	}
	test = p.Validate()
	if test == noError {
		t.Errorf("Validation should give empty gender error message")
	}

	p = Person{
		Name:   "John",
		Gender: "Pria",
		Age:    24,
	}
	test = p.Validate()
	if test == noError {
		t.Errorf("Validation should give invalid gender error message")
	}

	p = Person{
		Name:   "John",
		Gender: "Male",
	}
	test = p.Validate()
	if test == noError {
		t.Errorf("Validation should give empty age error message")
	}

	p = Person{
		Name:   "John",
		Gender: "Male",
		Age:    -1,
	}
	test = p.Validate()
	if test == noError {
		t.Errorf("Validation should give minus age error message")
	}
}

func (e error) New(s string) error {
	err := error{message: s}
	return err
}

func (p Person) Validate() error {
	errors := error{}
	if p.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if p.Gender != "Male" && p.Gender != "Female" {
		return errors.New("Gender is either Male or Female")
	}

	if p.Age < 0 {
		return errors.New("There is no such thing as negative age")
	}

	return errors
}

func main() {
	p := Person{
		Name:   "John",
		Gender: "Male",
		Age:    24,
	}
	fmt.Println(p.Validate())
}
