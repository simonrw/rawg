package main

func main() {
	builder := NewBuilder()
	rootURL := "https://example.com"
	builder.CreateClient(rootURL)

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
	err := builder.GenerateTypes(apiTypes)
	if err != nil {
		panic(err)
	}
	builder.Save("/tmp/out.go")
}
