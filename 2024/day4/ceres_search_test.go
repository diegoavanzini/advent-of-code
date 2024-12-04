package day4

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMASSearch_WithBigInput(t *testing.T) {
	want := 2507
	input := ReadFile("input.txt")
	got := ceresSearch("XMAS", input)
	assert.Equal(t, want, got)
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func TestXMAS_Search(t *testing.T) {
	want := 18
	input := []string{
		`MMMSXXMASM`,
		`MSAMXMSMSA`,
		`AMXSXMAAMM`,
		`MSAMASMSMX`,
		`XMASAMXAMM`,
		`XXAMMXXAMA`,
		`SMSMSASXSS`,
		`SAXAMASAAA`,
		`MAMMMXMMMM`,
		`MXMXAXMASX`,
	}
	got := ceresSearch("XMAS", input)
	assert.Equal(t, want, got)
}

func TestXMAS_DummySearch(t *testing.T) {
	want := 2
	input := []string{
		`MMMSXXMASM`,
		`MSAMXMSMSA`}
	got := ceresSearch("XMAS", input)
	assert.Equal(t, want, got)
}
