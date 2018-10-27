package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateTypes generates types and prints to stdout
func GenerateTypes(types ApiTypes) error {
	for typeIdx, typedef := range types {
		argDefs := make([]jen.Code, len(typedef.Parameters))

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

		structName := fmt.Sprintf("Api%s", typedef.Name)
		stmt := jen.Type().Id(structName).Struct(argDefs...)
		fmt.Printf("%#v", stmt)
	}
	return nil
}
