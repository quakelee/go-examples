package models

import (
	"reflect"
)

// Test is a sample struct to define an empty check method
type Test struct {

	// ID
	ID int `json:"id,omitempty"`

	// Name

	// type
	Type *string `json:"type"`

	// Status
	Status bool `json:"status,omitempty"`
}

// IsEmpty check struct is empty

/*
// This is object version
func (m Test) IsEmpty() bool {
	return reflect.DeepEqual(m, Test{})
}
*/

// IsEmpty check struct is empty
// This is pointer version
func (m *Test) IsEmpty() bool {
	return reflect.DeepEqual(*m, Test{})
}
