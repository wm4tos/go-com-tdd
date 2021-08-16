package main

import "testing"

func TestHello(t *testing.T) {
	name := "Wes"
	result := Hello("Wes")
	expected := "Hello, " + name + "!"

	if result != expected {
		t.Errorf("result '%s' | expected '%s'", result, expected)
	}
}
