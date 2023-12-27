package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("hello there are you sure\n")
	exp := 5
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}

}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("words are cool\n so are you\n thanks")
	exp := 3
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("Are we sure we're safe\n")
	exp := 5
	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}
