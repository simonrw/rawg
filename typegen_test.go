package main

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
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

	file := jen.NewFile("main")
	err := GenerateTypes(file, apiTypes)
	if err != nil {
		t.Errorf("%s", err)
	}

	text := stripTabs(fmt.Sprintf("%#v", file))
	if text != expected {
		t.Errorf("Result:\n%#v\n!=\n%#v", text, expected)
	}
}
