package main

import "testing"

func TestEndpointGen(t *testing.T) {
	t.Skip("Not implemented")

	endpoint := Endpoint{
		Name:    "records",
		Url:     "/records",
		Method:  "GET",
		Returns: "Suggestion",
	}

	expected := `func FetchApiRecords() (*Suggestion, error) {
	}
	`

	text, err := GenerateEndpoint(endpoint)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if text != expected {
		t.Errorf("Result:\n%s != expected:\n%s", text, expected)
	}
}
