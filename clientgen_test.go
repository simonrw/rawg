package main

import "testing"

func TestCreateClient(t *testing.T) {
	rootURL := "http://httpbin.org"
	stmts := CreateClient(rootURL)

	expected := `type Client struct {
	rootURL string
}

func NewClient() Client {
	return Client{rootURL: "http://httpbin.org"}
}`

	if stmts != expected {
		t.Errorf("Result:\n%#v\n!=\n%#v", stmts, expected)
	}
}
