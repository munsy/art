package main

import (
	"testing"
)

func TestNewClientNil(t *testing.T) {
	c := NewClient("something")

	if nil == c {	
		t.Fatal("nil wasn't expected but was returned")
	}
}
