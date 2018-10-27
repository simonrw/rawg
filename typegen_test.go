package main

import (
	"fmt"
	"testing"
)

// Helper function for performing the tests
func compareGeneratedTypes(types ApiTypes, expected string) error {
	text, err := GenerateTypes(types)
	if err != nil {
		return fmt.Errorf("Unexpected error: %v", err)
	}

	if text != expected {
		return fmt.Errorf("Result: \n%s != expected:\n%s", text, expected)
	}

	return nil
}

func TestTypegen(t *testing.T) {
	apiTypes := ApiTypes{
		ApiType{
			Name: "Record",
			Parameters: TypeParameters{
				TypeParameter{
					Name: "foo",
					Type: "int",
				},
			},
		},
	}

	expected := `type ApiRecord struct {
	foo int
}
`
	err := compareGeneratedTypes(apiTypes, expected)
	if err != nil {
		t.Errorf("%s", err)
	}
}
