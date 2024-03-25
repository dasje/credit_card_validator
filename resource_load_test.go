package main

import (
	"testing"
)
func TestResourceLoad(t *testing.T) {
	a,b,c := loadResources()
	if a == nil {
		t.Error("Industry resources not loaded.")
	}
	if b == nil {
		t.Error("Issuer resources not loaded.")
	}
	if c == nil {
		t.Error("Card regex resources not loaded.")
	}
}

// func TestRegex(t *testing.T) {
// 	a,b,c := loadResources()
// }