package day2

import "math"

func safeReports(reports [][]int) (result int) {
	safeCounter := 0
	for iReport := 0; iReport < len(reports); iReport++ {
		currentReport := reports[iReport]
		numberOfLevels := len(currentReport)
		var prevLevel = currentReport[0]
		for asc, desc, iLevel := 1, 1, 1; iLevel < numberOfLevels; iLevel++ {
			currentLevel := currentReport[iLevel]
			if math.Abs(float64(currentLevel-prevLevel)) <= 3 &&
				math.Abs(float64(currentLevel-prevLevel)) > 0 {
				if currentLevel > prevLevel {
					asc++
				} else {
					desc++
				}
			}
			prevLevel = currentLevel
			if asc == numberOfLevels || desc == numberOfLevels {
				safeCounter++
			}
		}
	}
	return safeCounter
}

func safeReportsP2(reports [][]int) (result int) {
	safeCounter := 0
	for iReport := 0; iReport < len(reports); iReport++ {
		if isSafe(reports[iReport], true) {
			safeCounter++
		}
	}
	return safeCounter
}

func isSafe(currentReport []int, retry bool) bool {
	var directionAsc bool
	var prevLevel = currentReport[0]
	for iLevel := 1; iLevel < len(currentReport); iLevel++ {
		currentLevel := currentReport[iLevel]
		if iLevel == 1 {
			directionAsc = currentLevel > prevLevel
		}
		if math.Abs(float64(currentLevel-prevLevel)) > 3 ||
			directionAsc != (currentLevel > prevLevel) {
			if !retry {
				return false
			}
			correctReport := append(currentReport[:iLevel], currentReport[iLevel+1:]...)
			return isSafe(correctReport, false)
		}
		prevLevel = currentLevel
	}
	return true
}
