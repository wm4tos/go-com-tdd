package main

import "testing"

func TestHello(t *testing.T) {
	name := "Wes"

	verifyMessage := func(t *testing.T, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("result '%s' | expected '%s'", result, expected)
		}
	}

	t.Run("Should say 'Hello, Wes!'", func(t *testing.T) {
		result := Hello("Wes")
		expected := "Hello, " + name + "!"

		verifyMessage(t, result, expected)
	})

	t.Run("Should say 'Hello, world!' for empty string", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world!"

		verifyMessage(t, result, expected)
	})
}
