package day4

import (
	"fmt"
	"math"
)

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

func (current *Location) IsMas(input []string) bool {
	// first and last rows and first or last columns are not valid
	if current.Row == 0 || current.Col == 0 ||
		current.Row == len(input)-1 || current.Col == len(input[0])-1 {
		return false
	}
	up_left := string(input[current.Row-1][current.Col-1])
	down_right := string(input[current.Row+1][current.Col+1])
	down_left := string(input[current.Row+1][current.Col-1])
	up_right := string(input[current.Row-1][current.Col+1])
	return (up_left == "M" && down_right == "S" ||
		up_left == "S" && down_right == "M") &&
		(down_left == "M" && up_right == "S" ||
			down_left == "S" && up_right == "M")
}
