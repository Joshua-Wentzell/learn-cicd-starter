package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	testHeader := http.Header{
		"Authorization": []string{""},
	}
	result, _ := GetAPIKey(testHeader)
	if result != "" {
		t.Fatalf("expected result to be empty")
	}
}
