package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := add(10, 20)
	if a != 300 {
		t.Error("add(10, 20) should be equal to 300")
	}
}

func TestAdd2(t *testing.T) {
	a := add(10, 20)
	if a != 30 {
		t.Error("add(10, 20) should be equal to 30")
	}
}
