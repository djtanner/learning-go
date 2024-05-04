package main

import (
	"testing"
)

func TestContainsZ(t *testing.T) {
	t.Run("testing contains z", func(t *testing.T) {
		got := ContainsZ("Zebra")
		want := true
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
