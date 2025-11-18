package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!', when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("クリスさん", "にほんじん")
		want := "クリスさん, こんにちは!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese, no name", func(t *testing.T) {
		got := Hello("", "にほんじん")
		want := "World, こんにちは!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Chris", "italiano")
		want := "Boungiorno, Chris!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	// tell the test suite that this is a helper function
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
