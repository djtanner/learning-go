package main

import (
	"reflect"
	"testing"
)

func TestContainsZ(t *testing.T) {
	t.Run("testing contains z", func(t *testing.T) {
		got := ContainsZ("Zebra")
		want := true
		assertCorrectMessage(t, got, want)
	})
}

func TestSplitting(t *testing.T) {
	t.Run("testing splitting", func(t *testing.T) {
		got := Splitting("Hello World")
		want := []string{"Hello", "World"}

		assertSplitMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func assertSplitMessage(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
