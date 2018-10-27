package main

import "testing"

func TestTypegen(t *testing.T) {
	apiTypes := []ApiType{
		ApiType{
			Name: "Record",
			Parameters: []TypeParameter{
				TypeParameter{
					Name: "foo",
					Type: "int",
				},
			},
		},
	}
	text, err := GenerateTypes(apiTypes)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := `type ApiRecord struct {
	foo int
}
`
	if text != expected {
		t.Errorf("Result: \n%s != expected:\n%s", text, expected)
	}
}
