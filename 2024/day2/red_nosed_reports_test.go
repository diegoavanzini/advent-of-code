package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeReports(t *testing.T) {
	want := 2
	var reports = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	got := safeReports(reports)
	assert.Equal(t, want, got)
}

func TestSafeReportsAoCInput(t *testing.T) {
	want := 483
	reports := ReadFile("input.txt")
	got := safeReports(reports)
	assert.Equal(t, want, got)
}

func TestSafeReportsPart2(t *testing.T) {
	want := 5
	var reports = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{1, 6, 4, 7, 9},
	}
	got := safeReportsP2(reports)
	assert.Equal(t, want, got)
}

func TestSafeReportsPart2WithAoCInput(t *testing.T) {
	want := 548
	reports := ReadFile("input.txt")
	got := safeReportsP2(reports)
	assert.Equal(t, want, got)
}

func ReadFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]int
	for scanner.Scan() {
		report := []int{}
		for _, v := range strings.Split(scanner.Text(), " ") {
			level, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, level)
		}
		lines = append(lines, report)
	}

	return lines
}
