package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
)

// TODO: work out why this is needed...
func stripTabs(text string) string {
	return strings.Replace(text, "\t", "", -1)
}

func renderStatementsToString(stmts *jen.Statement) string {
	text := fmt.Sprintf("%#v", stmts)
	return stripTabs(text)
}

func TestCreateClient(t *testing.T) {
	rootURL := "http://httpbin.org"
	stmt := jen.Empty()
	stmts := CreateClient(stmt, rootURL)

	expected := stripTabs(`type Client struct {
	rootURL string
}

func NewClient() Client {
	return Client{rootURL: "http://httpbin.org"}
}`)

	text := renderStatementsToString(stmts)

	if text != expected {
		t.Errorf("Result:\n%#v\n!=\n%#v", text, expected)
	}
}
