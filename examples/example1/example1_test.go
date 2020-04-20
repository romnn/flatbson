package main

import (
	"testing"
)

func TestCli(t *testing.T) {
	out := getLongestKey()
	expected := "metadata.city.name"
	if out != expected {
		t.Errorf("Got %s but expected %s", out, expected)
	}
}
