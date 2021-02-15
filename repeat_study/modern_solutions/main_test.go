package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	result := reverseString("abc")
	if result != "cba" {
		t.Errorf("expected cba instead got %s", result)
	}
}
