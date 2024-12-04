package day4

import (
	"fmt"
	"math"
)

func ceresSearch(word string, input []string) (found int) {
	// xLocations := searchFor(word[0], input)
	// mLocations := searchFor(word[1], input)
	// aLocations := searchFor(word[2], input)
	// sLocations := searchFor(word[3], input)

	var locations = make([][]Location, len(word))
	for i := 0; i < len(word); i++ {
		locations[i] = searchFor(word[i], input)
	}

	return searchWord(locations)

}

func searchWord(locations [][]Location) (found int) {
	xLocations := locations[0]
	mLocations := locations[1]

	for iX := 0; iX < len(xLocations); iX++ {
		for iM := 0; iM < len(mLocations); iM++ {
			if dir, ok := xLocations[iX].Near(mLocations[iM]); ok {
				found += searchNear(locations[2:], mLocations[iM], dir)
			}
		}
	}
	return
}

func searchNear(currentCharLocations [][]Location, current Location, dir Direction) (found int) {
	if len(currentCharLocations) == 0 {
		return 1
	}
	for i := 0; i < len(currentCharLocations[0]); i++ {
		if newDir, ok := current.Near(currentCharLocations[0][i]); ok &&
			newDir == dir {
			found += searchNear(currentCharLocations[1:], currentCharLocations[0][i], dir)
			return
		}
	}
	return
}

type Location struct {
	Row int
	Col int
}

type Direction int

const (
	LEFT_DIRECTION Direction = iota
	RIGHT_DIRECTION
	UP_DIRECTION
	DOWN_DIRECTION
	TRAVERSAL_RIGHT_DOWN_DIRECTION
	TRAVERSAL_RIGHT_UP_DIRECTION
	TRAVERSAL_LEFT_DOWN_DIRECTION
	TRAVERSAL_LEFT_UP_DIRECTION
)

func (l1 *Location) String() string {
	return fmt.Sprintf("[%d:%d]", l1.Row, l1.Col)
}

func (current *Location) Near(next Location) (Direction, bool) {
	var resultDirection Direction
	near := (math.Abs(float64(next.Col-current.Col)) == 1 ||
		math.Abs(float64(next.Col-current.Col)) == 0) &&
		(math.Abs(float64(next.Row-current.Row)) == 1 ||
			math.Abs(float64(next.Row-current.Row)) == 0)
	if near {
		if next.Row == current.Row &&
			next.Col-current.Col == 1 {
			return RIGHT_DIRECTION, near
		}
		if next.Row == current.Row &&
			next.Col-current.Col == -11 {
			return LEFT_DIRECTION, near
		}
		if next.Col == current.Col &&
			next.Row-current.Row == 1 {
			return DOWN_DIRECTION, near
		}
		if next.Col == current.Col &&
			next.Row-current.Row == -1 {
			return UP_DIRECTION, near
		}
		if next.Row-current.Row == 1 &&
			next.Col-current.Col == 1 {
			return TRAVERSAL_RIGHT_DOWN_DIRECTION, near
		}
		if next.Row-current.Row == -1 &&
			next.Col-current.Col == 1 {
			return TRAVERSAL_RIGHT_UP_DIRECTION, near
		}
		if next.Row-current.Row == 1 &&
			next.Col-current.Col == -1 {
			return TRAVERSAL_LEFT_DOWN_DIRECTION, near
		}
		if next.Row-current.Row == -1 &&
			next.Col-current.Col == -1 {
			return TRAVERSAL_LEFT_UP_DIRECTION, near
		}
	}
	return resultDirection, near
}

func searchFor(b byte, input []string) (xLocations []Location) {
	for iRow := 0; iRow < len(input); iRow++ {
		for iCol := 0; iCol < len(input[iRow]); iCol++ {
			if input[iRow][iCol] == b {
				xLocations = append(xLocations, Location{Row: iRow, Col: iCol})
			}
		}
	}
	return
}
