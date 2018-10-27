package main

import (
	"bytes"

	"github.com/dave/jennifer/jen"
)

func CreateClient(rootURL string) string {
	b := &bytes.Buffer{}

	stmt := jen.Type().Id("Client").Struct(
		jen.Id("rootURL").String(),
	)
	stmt = stmt.Add(jen.Line())
	stmt = stmt.Add(jen.Func().Id("NewClient").Params().Id("Client").Block(
		jen.Return(
			jen.Id("Client").Values(jen.Dict{
				jen.Id("rootURL"): jen.Lit(rootURL),
			}),
		),
	),
	)
	err := stmt.Render(b)
	if err != nil {
		panic(err)
	}

	return b.String()
}
