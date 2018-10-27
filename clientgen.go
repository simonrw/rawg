package main

import (
	"github.com/dave/jennifer/jen"
)

func (b *Builder) CreateClient(rootURL string) {
	b.file.Type().Id("Client").Struct(
		jen.Id("rootURL").String(),
	)

	b.file.Line()
	b.file.Func().Id("NewClient").Params().Id("Client").Block(
		jen.Return(
			jen.Id("Client").Values(jen.Dict{
				jen.Id("rootURL"): jen.Lit(rootURL),
			}),
		),
	)
}
