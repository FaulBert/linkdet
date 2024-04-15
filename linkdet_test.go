package main

import (
	"testing"
)

func TestDetective(t *testing.T) {
	short_url := "http://bit.ly/abc123"
	url, err := detective(short_url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if url == "" {
		t.Errorf("Expected URL, got empty string")
	}
}
