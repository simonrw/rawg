package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Definition struct {
	Endpoints Endpoints `yaml:"endpoints,flow"`
	Types     ApiTypes  `yaml:"types,flow"`
}

type Endpoint struct {
	Url    string `yaml:"url"`
	Method string `yaml:"method"`
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

func ParseDefinitionFromFile(filename string) (*Definition, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var definition Definition
	err = yaml.Unmarshal(bytes, &definition)
	if err != nil {
		return nil, err
	}

	return &definition, nil
}
