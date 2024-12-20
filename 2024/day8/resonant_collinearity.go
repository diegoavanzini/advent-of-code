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

func ResonantCollinearity(input []string, part2 ...bool) int {
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
				anps := getAntinodesPosition(ap[an1], ap[an2], xLimit, yLimit, part2...)
				if len(part2) > 0 && part2[0] && len(anps) > 0 {
					antinodesCounterResult[fmt.Sprintf("%d-%d", ap[an1].X, ap[an1].Y)] = ap[an1]
					antinodesCounterResult[fmt.Sprintf("%d-%d", ap[an2].X, ap[an2].Y)] = ap[an2]
				}
				for i := 0; i < len(anps); i++ {
					output[anps[i].Y] = output[anps[i].Y][:anps[i].X] + "#" + output[anps[i].Y][anps[i].X+1:]
					antinodesCounterResult[fmt.Sprintf("%d-%d", anps[i].X, anps[i].Y)] = anps[i]
				}
			}
		}
	}
	Visualize(output, len(antinodesCounterResult))
	return len(antinodesCounterResult)
}

func Visualize(output []string, len int) {
	for _, o := range output {
		fmt.Printf("%s\n", o)
	}
	fmt.Printf("counter: %d\n\n", len)
}

func intoTheArea(anp AntennaPosition, xLimit int, yLimit int) bool {
	return anp.X <= xLimit &&
		anp.Y <= yLimit &&
		anp.X >= 0 &&
		anp.Y >= 0
}

func getAntinodesPosition(antennaPosition1, antennaPosition2 AntennaPosition, xLimit int, yLimit int, part2 ...bool) (anp []AntennaPosition) {

	xDistance := antennaPosition1.X - antennaPosition2.X
	yDistance := antennaPosition1.Y - antennaPosition2.Y
	var i int
	for {
		i++
		antinode1 := AntennaPosition{Value: "#"}
		antinode2 := AntennaPosition{Value: "#"}
		if xDistance >= 0 {
			antinode1.X = antennaPosition1.X + (xDistance * i)
			antinode2.X = antennaPosition2.X - (xDistance * i)
		} else {
			antinode1.X = antennaPosition1.X + (xDistance * i)
			antinode2.X = antennaPosition2.X - (xDistance * i)
		}
		if yDistance >= 0 {
			antinode1.Y = antennaPosition2.Y - (yDistance * i)
			antinode2.Y = antennaPosition1.Y + (yDistance * i)
		} else {
			antinode1.Y = antennaPosition1.Y + (yDistance * i)
			antinode2.Y = antennaPosition2.Y - (yDistance * i)
		}
		if !intoTheArea(antinode1, xLimit, yLimit) && !intoTheArea(antinode2, xLimit, yLimit) {
			return
		}
		if intoTheArea(antinode1, xLimit, yLimit) {
			anp = append(anp, antinode1)
		}
		if intoTheArea(antinode2, xLimit, yLimit) {
			anp = append(anp, antinode2)
		}
		if len(part2) == 0 || !part2[0] {
			return
		}
	}
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
