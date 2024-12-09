package day6

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHowManyDistinctPositionWithAoCInput(t *testing.T) {
	want := 4374
	m := ReadInputFile("./input.txt")
	situationMap := NewSituationMap(m)
	got := situationMap.HowManyDistinctPosition()
	assert.Equal(t, want, got)
}

func TestHowManyDistinctPosition(t *testing.T) {
	want := 41
	m := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	situationMap := NewSituationMap(m)
	got := situationMap.HowManyDistinctPosition()
	assert.Equal(t, want, got)
}
func ReadInputFile(filename string) []string {
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
