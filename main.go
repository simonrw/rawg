package main

import (
	"fmt"
	"log"

	. "github.com/dave/jennifer/jen"
	flag "github.com/spf13/pflag"
)

func main() {
	var port *int = flag.IntP("port", "p", 8080, "Port to listen on")
	flag.Parse()

	_ = port

	def, err := ParseDefinitionFromFile("definition.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", def)

	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello world")),
	)
}
