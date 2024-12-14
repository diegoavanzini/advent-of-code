package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTotalCalibrationResultAoCInput(t *testing.T) {
	want := 483
	parsedInput := ReadFile("input.txt")
	got := getTotalCalibrationResult(parsedInput)
	assert.Equal(t, want, got)
}

func TestFindTotalCalibrationResult(t *testing.T) {
	var want = 3749
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	parsedInput, err := parse(input)
	if err != nil {
		t.Fatal(err)
	}
	got := getTotalCalibrationResult(parsedInput)
	assert.Equal(t, want, got)
}

func parse(input string) ([]CalibrationResult, error) {
	result := []CalibrationResult{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		currentCalibration, err := parseRow(row)
		if err != nil {
			return result, err
		}
		result = append(result, currentCalibration)
	}
	return result, nil
}

func parseRow(row string) (CalibrationResult, error) {
	splittedRow := strings.Split(row, ":")
	rowResult, err := strconv.Atoi(splittedRow[0])
	if err != nil {
		return CalibrationResult{}, err
	}
	calibrations := strings.Split(strings.Trim(splittedRow[1], " "), " ")
	calibrationsInt := make([]int, len(calibrations))
	for i, v := range calibrations {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			return CalibrationResult{}, err
		}
		calibrationsInt[i] = vInt
	}
	currentCalibration := CalibrationResult{
		Result:       rowResult,
		Calibrations: calibrationsInt,
	}
	return currentCalibration, nil
}

func ReadFile(filename string) []CalibrationResult {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []CalibrationResult
	for scanner.Scan() {
		report, err := parseRow(scanner.Text())
		// for _, v := range strings.Split(scanner.Text(), " ") {
		// 	level, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		// 	report = append(report, level)
		// }
		lines = append(lines, report)
	}
	return lines
}
