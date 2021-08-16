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
		result := Hello(name, "")
		expected := "Hello, " + name + "!"

		verifyMessage(t, result, expected)
	})

	t.Run("Should say 'Hello, world!' for empty string", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world!"

		verifyMessage(t, result, expected)
	})

	t.Run("Should print 'Olá, Wes!' because language is portuguese", func(t *testing.T) {
		result := Hello(name, "portuguese")
		expected := "Olá, " + name + "!"

		verifyMessage(t, result, expected)
	})

	t.Run("Should print 'Bonjour, Wes!' because language is french", func(t *testing.T) {
		result := Hello(name, "french")
		expected := "Bonjour, " + name + "!"

		verifyMessage(t, result, expected)
	})
}
