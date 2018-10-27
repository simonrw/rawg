package main

import "github.com/dave/jennifer/jen"

func main() {
	file := jen.NewFile("main")
	rootURL := "https://example.com"
	CreateClient(file, rootURL)

	apiTypes := ApiTypes{
		ApiType{
			Name: "Record",
			Parameters: TypeParameters{
				TypeParameter{
					Name: "foo",
					Type: "int",
				},
			},
		},
	}
	err := GenerateTypes(file, apiTypes)
	if err != nil {
		panic(err)
	}

	file.Save("/tmp/out.go")
}
