package main

import (
	"fmt"

	"github.com/quakelee/go-examples/commonfunctions/libs"
	"github.com/quakelee/go-examples/commonfunctions/models"
)

func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.
	libs.Sum(1, 2)
	libs.Sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	libs.Sum(nums...)

	// Covert int Slice or Array to String
	fmt.Println(libs.IntSliceToString(nums, ", "))

	// Example of show how to check struct empty
	var test models.Test
	//tt := "empty test"
	//test.Type = &tt
	pt := &test
	fmt.Println(pt.IsEmpty())
}
