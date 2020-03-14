package sampleapp

import (
	"testing"
)

func TestAdd(t *testing.T) {
	actual := add(10, 20)
	expected := 30
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestReadFile(t *testing.T) {
	actual := readFile("input")
	expected := "aaa\r\nbbb\r\nccc\r\n"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
