package main

import (
	"github.com/dave/jennifer/jen"
)

func CreateClient(file *jen.File, rootURL string) {
	file.Type().Id("Client").Struct(
		jen.Id("rootURL").String(),
	)

	file.Line()
	file.Func().Id("NewClient").Params().Id("Client").Block(
		jen.Return(
			jen.Id("Client").Values(jen.Dict{
				jen.Id("rootURL"): jen.Lit(rootURL),
			}),
		),
	)
}
