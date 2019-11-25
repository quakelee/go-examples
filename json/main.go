package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type sweetData struct {
	Foo string `json:"foo"`
}
type sweetDataNullable struct {
	Foo *string `json:"foo"`
}

// This is a example to show how nullable field in struct work with json marshal and unmarshal
func main() {
	var s1 sweetData
	var s2 sweetDataNullable

	if err := json.NewDecoder(strings.NewReader("{}")).Decode(&s1); err != nil {
		panic(err)
	}
	fmt.Printf("s1 (non-nullable): %#v\n", s1)

	if err := json.NewDecoder(strings.NewReader("{}")).Decode(&s2); err != nil {
		panic(err)
	}
	fmt.Printf("s2 (nullable): %#v\n", s2)

	s3, err := json.Marshal(&s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("s3 %#v\n", string(s3))

	s4, err := json.Marshal(&s2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("s4 %#v\n", string(s4))
}