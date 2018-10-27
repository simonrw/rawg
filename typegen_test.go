package main

import (
	"fmt"
	"testing"
)

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

	expected := stripTabs(`package main

type ApiRecord struct {
	foo int
}
`)

	builder := NewBuilder()
	err := builder.GenerateTypes(apiTypes)
	if err != nil {
		t.Errorf("%s", err)
	}

	text := stripTabs(fmt.Sprintf("%#v", builder.file))
	if text != expected {
		t.Errorf("Result:\n%#v\n!=\n%#v", text, expected)
	}
}
