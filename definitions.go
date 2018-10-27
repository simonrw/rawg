package main

type Definition struct {
	Endpoints Endpoints `yaml:"endpoints,flow"`
	Types     ApiTypes  `yaml:"types,flow"`
}

type Endpoint struct {
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
	Method  string `yaml:"method"`
	Returns string `yaml:"returns"`
}

type ApiType struct {
	Name       string         `yaml:"name"`
	Parameters TypeParameters `yaml:"parameters"`
}

type TypeParameter struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Endpoints []Endpoint
type ApiTypes []ApiType
type TypeParameters []TypeParameter
