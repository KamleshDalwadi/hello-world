package main

import "testing"

// Test greet function
func TestGreet(t *testing.T) {
	str, err := greet("there")
	expected := "Hi there"
	if err != nil {
		t.Errorf("Error greeting %v", err)
	}
	if str != expected {
		t.Errorf("Expected %s, but got %s", expected, str)
	}
}
