package main

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGETapiKey(t *testing.T) {
	header := http.Header{}
	const apiKey string = "1234"
	header.Set("Authorization", "ApiKey"+" "+apiKey)
	result, err := auth.GetAPIKey(header)
	if err != nil {
		t.Errorf("Unexpected Error:%v", err)
	}
	if result != apiKey {
		t.Errorf("Key was not extracted correctly: Expected %v, Got %v", apiKey, result)
	}
}
func TestGETapiKeyWrong(t *testing.T) {
	header := http.Header{}
	const apiKey string = "1234"
	header.Set("Authorization", "ApisKey"+" "+apiKey)
	result, err := auth.GetAPIKey(header)
	if err != nil {
		return
	}
	if result == apiKey {
		t.Errorf("Key should have not been extracted")
	}
}
