package domain

import (
	"bytes"
	"encoding/gob"
)

type Employee struct {
	Name   string
	Office Address
}

// DeepCopy function using serialization
func (e *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	_ = encoder.Encode(e)

	decoder := gob.NewDecoder(&b)
	result := Employee{}
	_ = decoder.Decode(&result)
	return &result
}

// Factories function
var mainOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		StreetAddress: "ABC Street",
		City:          "Jakarta",
	},
}

var secondOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		StreetAddress: "ZZZ Street",
		City:          "Bandung",
	},
}

// Constructor
func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&secondOffice, name, suite)
}

func newEmployee(employee *Employee, name string, suite int) *Employee {
	result := employee.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}
