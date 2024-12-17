package day8

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResonantCollinearityAoCInput(t *testing.T) {
	want := 336
	input := ReadInputFile("./input.txt")

	got := ResonantCollinearity(input)
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

func TestResonantCollinearity(t *testing.T) {
	want := 14
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	got := ResonantCollinearity(input)
	assert.Equal(t, want, got)
}

func TestGetAntennaPositionsGroupedByValue(t *testing.T) {
	input := []string{
		"............",
		".....0......",
		".......0....",
		"............",
	}
	ap := getAntennaPositionsGroupedByValue(input)
	assert.Equal(t, 1, len(ap))
	assert.Equal(t, 2, len(ap["0"]))
}

func TestGetAntinodesPosition(t *testing.T) {
	ap1 := AntennaPosition{
		Value: "0",
		X:     5,
		Y:     1,
	}
	ap2 := AntennaPosition{
		Value: "0",
		X:     7,
		Y:     2,
	}
	an := getAntinodesPosition(ap1, ap2)
	assert.Equal(t, 2, len(an))
	assert.Equal(t, 3, an[0].X)
	assert.Equal(t, 0, an[0].Y)
	assert.Equal(t, 9, an[1].X)
	assert.Equal(t, 3, an[1].Y)
}

func TestRegularExpression(t *testing.T) {
	row := ".....0...0..."
	r := regexp.MustCompile("[a-zA-Z0-9]")
	matchAntennas := r.FindAllStringIndex(row, -1)
	initialIndex := matchAntennas[0][0]
	finalIndex := matchAntennas[0][1]
	assert.Equal(t, 5, initialIndex)
	assert.Equal(t, 6, finalIndex)
}
