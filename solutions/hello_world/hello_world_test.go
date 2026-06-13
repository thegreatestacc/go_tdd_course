package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello(english, "Vladimir")
		expected := "Hello Vladimir"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		actual := Hello(english, "")
		expected := "Hello World"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Spanish", "Elodie")
		expected := "Hola Elodie"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in French", func(t *testing.T) {
		actual := Hello("French", "Macron")
		expected := "Bonjour Macron"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("unknown language", func(t *testing.T) {
		actual := Hello("***", "Unknown name")
		expected := "Hello Unknown name"

		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual is '%s', expected is '%s'",
			actual, expected)
	}
}
