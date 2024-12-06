package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintQueueWithAoCInput_Adjusted(t *testing.T) {
	want := 6456
	rules := ReadRulesFile("./rules.txt")
	input := ReadInputFile("./input.txt")
	p := NewPrinter(rules)
	got := p.PrintQueue(input, true)
	assert.Equal(t, want, got)
}
func TestPrintQueue_Adjusted(t *testing.T) {
	want := 123
	rules := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}
	input := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	p := NewPrinter(rules)
	got := p.PrintQueue(input, true)
	assert.Equal(t, want, got)
}

func TestPrintQueueWithAoCInput(t *testing.T) {
	want := 4569
	rules := ReadRulesFile("./rules.txt")
	input := ReadInputFile("./input.txt")
	p := NewPrinter(rules)
	got := p.PrintQueue(input, false)
	assert.Equal(t, want, got)
}
func TestPrintQueue(t *testing.T) {
	want := 143
	rules := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}
	input := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	p := NewPrinter(rules)
	got := p.PrintQueue(input, false)
	assert.Equal(t, want, got)
}

func ReadRulesFile(filename string) []string {
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

func ReadInputFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]int
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		splitLineInt := make([]int, len(splitLine))
		for i, v := range splitLine {
			splitLineInt[i], _ = strconv.Atoi(v)
		}
		lines = append(lines, splitLineInt)
	}

	return lines
}
