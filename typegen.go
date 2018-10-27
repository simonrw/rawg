package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// GenerateTypes generates types and prints to stdout
func GenerateTypes(types ApiTypes) (string, error) {

	// Writer for final result
	var b strings.Builder

	// Iterate over the types found in the definition
	for typeIdx, typedef := range types {
		argDefs := make([]jen.Code, len(typedef.Parameters))

		// Iterate over the parameter definitions
		// TODO: handle arrays and maps
		for paramIdx, param := range typedef.Parameters {
			switch param.Type {
			case "int":
				argDefs[paramIdx] = jen.Id(param.Name).Int()
			case "float32":
				argDefs[paramIdx] = jen.Id(param.Name).Float32()
			case "string":
				argDefs[paramIdx] = jen.Id(param.Name).String()
			default:
				return "", fmt.Errorf("Type %d, parameter %d, invalid type: %s", typeIdx, paramIdx, param.Type)
			}
		}

		// The struct name should not clash with other structs in the user's system, so
		// we prepend "Api" to the struct name
		structName := fmt.Sprintf("Api%s", typedef.Name)

		// Build the actual struct definition from the argument definitions and add
		// to the string builder
		stmt := jen.Type().Id(structName).Struct(argDefs...)
		fmt.Fprintf(&b, "%#v\n", stmt)
	}

	// Return the final generated code
	return b.String(), nil
}
