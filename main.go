package main

import (
	"fmt"
	"io/ioutil"
	"log"

	. "github.com/dave/jennifer/jen"
	flag "github.com/spf13/pflag"
	yaml "gopkg.in/yaml.v2"
)

type Definition struct {
}

func main() {
	var port *int = flag.IntP("port", "p", 8080, "Port to listen on")
	flag.Parse()

	_ = port

	bytes, err := ioutil.ReadFile("definition.yml")
	if err != nil {
		log.Fatal(err)
	}

	var def Definition
	err = yaml.Unmarshal(bytes, &def)
	if err != nil {
		log.Fatal(err)
	}

	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello world")),
	)
	fmt.Printf("%#v", f)
}
