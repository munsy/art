package main

import (
	"fmt"
	"testing"
)

func TestNewClientNoUsername(t *testing.T) {
	_, err := NewClient("", "asdf")

	if nil == err {	
		t.Fatal("error was expected but not returned")
	}

	if err.Error() != "no username supplied" {
		t.Fatal(fmt.Sprintf("error mismatch, expected %s, got %s", "no username supplied", err.Error()))
	}
}

func TestNewClientNoPassword(t *testing.T) {
	_, err := NewClient("asdf", "")

	if nil == err {	
		t.Fatal("error was expected but not returned")
	}

	if err.Error() != "no password supplied" {
		t.Fatal(fmt.Sprintf("error mismatch, expected %s, got %s", "no password supplied", err.Error()))
	}
}
