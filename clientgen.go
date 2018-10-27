package main

import (
	"github.com/dave/jennifer/jen"
)

func CreateClient(stmt *jen.Statement, rootURL string) *jen.Statement {
	stmt = stmt.Add(jen.Type().Id("Client").Struct(
		jen.Id("rootURL").String(),
	))
	stmt = stmt.Add(jen.Line())
	stmt = stmt.Add(jen.Func().Id("NewClient").Params().Id("Client").Block(
		jen.Return(
			jen.Id("Client").Values(jen.Dict{
				jen.Id("rootURL"): jen.Lit(rootURL),
			}),
		),
	),
	)
	return stmt
}
