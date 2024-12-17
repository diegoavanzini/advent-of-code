package day8

import (
	"fmt"
	"regexp"
)

type AntennaPosition struct {
	Value string
	X     int
	Y     int
}

func ResonantCollinearity(input []string) int {
	output := input
	var antinodesCounterResult = map[string]AntennaPosition{}
	yLimit := len(input) - 1
	xLimit := len(input[0]) - 1
	aps := getAntennaPositionsGroupedByValue(input)
	for _, ap := range aps {
		if len(ap) < 2 {
			continue
		}
		for an1 := 0; an1 < len(ap); an1++ {
			for an2 := an1; an2 < len(ap); an2++ {
				if an1 == an2 {
					continue
				}
				anp := getAntinodesPosition(ap[an1], ap[an2])
				if intoTheArea(anp[0], xLimit, yLimit) {
					output[anp[0].Y] = output[anp[0].Y][:anp[0].X] + "#" + output[anp[0].Y][anp[0].X+1:]
					antinodesCounterResult[fmt.Sprintf("%d%d", anp[0].X, anp[0].Y)] = anp[0]
				}
				if intoTheArea(anp[1], xLimit, yLimit) {
					output[anp[1].Y] = output[anp[1].Y][:anp[1].X] + "#" + output[anp[1].Y][anp[1].X+1:]
					antinodesCounterResult[fmt.Sprintf("%d%d", anp[1].X, anp[1].Y)] = anp[1]
				}
			}
		}
	}
	return len(antinodesCounterResult)
}

func intoTheArea(anp AntennaPosition, xLimit int, yLimit int) bool {
	return anp.X <= xLimit &&
		anp.Y <= yLimit &&
		anp.X >= 0 &&
		anp.Y >= 0
}

func getAntinodesPosition(antennaPosition1, antennaPosition2 AntennaPosition) []AntennaPosition {
	antinode1 := AntennaPosition{Value: "#"}
	antinode2 := AntennaPosition{Value: "#"}
	xDistance := antennaPosition1.X - antennaPosition2.X
	if xDistance >= 0 {
		antinode1.X = antennaPosition1.X + xDistance
		antinode2.X = antennaPosition2.X - xDistance
	} else {
		antinode1.X = antennaPosition1.X + xDistance
		antinode2.X = antennaPosition2.X - xDistance
	}
	yDistance := antennaPosition1.Y - antennaPosition2.Y
	if yDistance >= 0 {
		antinode1.Y = antennaPosition2.Y - yDistance
		antinode2.Y = antennaPosition1.Y + yDistance
	} else {
		antinode1.Y = antennaPosition1.Y + yDistance
		antinode2.Y = antennaPosition2.Y - yDistance
	}
	return []AntennaPosition{antinode1, antinode2}
}

func getAntennaPositionsGroupedByValue(input []string) map[string][]AntennaPosition {
	antennaPositions := map[string][]AntennaPosition{}
	for i := 0; i < len(input); i++ {
		row := input[i]
		r := regexp.MustCompile("[a-zA-Z0-9]")
		matchAntennas := r.FindAllStringIndex(row, -1)
		if matchAntennas == nil {
			continue
		}
		for _, antenna := range matchAntennas {
			firstIndex := antenna[0]
			secondIndex := antenna[1]
			value := row[firstIndex:secondIndex]
			ap := AntennaPosition{
				Value: value,
				Y:     i,
				X:     firstIndex,
			}
			antennaPositions[value] = append(antennaPositions[value], ap)
		}
	}
	return antennaPositions
}
