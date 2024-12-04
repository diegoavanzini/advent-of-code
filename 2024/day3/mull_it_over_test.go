package main

import (
	"log"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestMullItOverP2WithAoCInput(t *testing.T) {
	want := 82857512
	input := ReadFile("./input.txt")
	got := MullItOverP2(input)
	assert.Equal(t, want, got)
}

func TestMullItOverP2(t *testing.T) {
	want := 48
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	got := MullItOverP2(input)
	assert.Equal(t, want, got)
}

func TestMullItOver(t *testing.T) {
	want := 161
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got := MullItOver(input)
	assert.Equal(t, want, got)
}

func TestMullItOverWithAoCInput(t *testing.T) {
	want := 190604937
	input := ReadFile("./input.txt")
	got := MullItOver(input)
	assert.Equal(t, want, got)
}

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
