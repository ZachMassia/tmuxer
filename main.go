package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// All components take no arguments and return a string and error.
type component func() (string, error)

// Each component will register it's main function with the
// components map in it's init() func.
var components = make(map[string]component)

// register adds a component func to the components map.
func register(id string, f component) { components[id] = f }

// getVal is a convenience function which reads an ACPI data file
// and returns it as a string with whitespace trimmed.
func getVal(file string) (string, error) {
	resp, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	// Convert to string and trim whitespace
	str := string(resp)
	return strings.TrimSpace(str), nil
}

func main() {
	comps := []string{"bat"} // TODO: Add support for command line flags.

	// Call each components function, and add it's output string to output.
	output := make([]string, 0, len(comps))
	for _, c := range comps {
		if compFunc, ok := components[c]; ok {
			if compOut, err := compFunc(); err == nil {
				output = append(output, compOut)
			}
		}
	}
	fmt.Println(strings.Join(output, " │ "))
}
