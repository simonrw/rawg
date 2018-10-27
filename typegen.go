package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateTypes generates types and prints to stdout
func (b *Builder) GenerateTypes(types ApiTypes) error {

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
				return fmt.Errorf("Type %d, parameter %d, invalid type: %s", typeIdx, paramIdx, param.Type)
			}
		}

		// The struct name should not clash with other structs in the user's system, so
		// we prepend "Api" to the struct name
		structName := fmt.Sprintf("Api%s", typedef.Name)

		// Build the actual struct definition from the argument definitions and add
		// to the string builder
		b.file.Type().Id(structName).Struct(argDefs...)
	}

	// Return the final generated code
	return nil
}
