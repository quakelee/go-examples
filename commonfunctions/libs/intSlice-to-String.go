package libs

import (
	"fmt"
	"strings"
)

// IntSliceToString is a example to convert int slice/array to a string
func IntSliceToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}