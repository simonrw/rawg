package main

import "github.com/dave/jennifer/jen"

type Builder struct {
	file *jen.File
}

func NewBuilder() Builder {
	return Builder{
		file: jen.NewFile("main"),
	}
}

func (b *Builder) Save(filename string) error {
	return b.file.Save(filename)
}
