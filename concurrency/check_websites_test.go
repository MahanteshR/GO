package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "what://furthers.gets" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://blog.gypsydave5.com",
		"what://furthers.gets",
	}

	want := map[string]bool{
		"https://google.com":          true,
		"https://blog.gypsydave5.com": true,
		"what://furthers.gets":        false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
