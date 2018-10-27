package main

import (
	"fmt"
	"strings"
	"testing"
)

// TODO: work out why this is needed...
func stripTabs(text string) string {
	return strings.Replace(text, "\t", "", -1)
}

func TestCreateClient(t *testing.T) {
	rootURL := "http://httpbin.org"
	builder := NewBuilder()
	builder.CreateClient(rootURL)

	expected := stripTabs(`package main
	
type Client struct {
	rootURL string
}

func NewClient() Client {
	return Client{rootURL: "http://httpbin.org"}
}
`)

	text := stripTabs(fmt.Sprintf("%#v", builder.file))

	if text != expected {
		t.Errorf("Result:\n%#v\n!=\n%#v", text, expected)
	}
}
