package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

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
