package day6

type Position struct {
	X, Y int
}

type SituationMap struct {
	MapPositions []string
	positions    int
}

func NewSituationMap(situationMap []string) SituationMap {
	return SituationMap{
		MapPositions: situationMap,
	}
}

type Direction int

const (
	UP_DIRECTION Direction = iota
	LEFT_DIRECTION
	RIGHT_DIRECTION
	DOWN_DIRECTION
)

func (m *SituationMap) HowManyDistinctPosition() (positions int) {
	var currentDirection = UP_DIRECTION
	var currentPosition = m.findStartPosition()
	if currentPosition == nil {
		return positions
	}
	m.positions = 1
	m.walk(*currentPosition, currentDirection)
	return m.positions
}

func (m *SituationMap) findStartPosition() *Position {
	for iRow, row := range m.MapPositions {
		for iCol := 0; iCol < len(row); iCol++ {
			if row[iCol] == '^' {
				return &Position{X: iCol, Y: iRow}
			}
		}
	}
	return nil
}

func (m *SituationMap) walk(currentPosition Position, currentDirection Direction) (Position, bool) {
	if m.MapPositions[currentPosition.Y][currentPosition.X] != 'X' {
		m.positions++
	}
	row := m.MapPositions[currentPosition.Y]
	m.MapPositions[currentPosition.Y] = row[:currentPosition.X] + "X" + row[currentPosition.X+1:]
	newPosition := currentPosition
	switch currentDirection {
	case UP_DIRECTION:
		newPosition.Y--
	case RIGHT_DIRECTION:
		newPosition.X++
	case LEFT_DIRECTION:
		newPosition.X--
	case DOWN_DIRECTION:
		newPosition.Y++
	}
	if m.isCollision(newPosition) {
		currentDirection = m.changeDirection(currentDirection)
		return m.walk(currentPosition, currentDirection)
	}
	if (newPosition.X == 0 && currentDirection == LEFT_DIRECTION) ||
		(newPosition.X == len(m.MapPositions[newPosition.Y])-1 && currentDirection == RIGHT_DIRECTION) ||
		(newPosition.Y == 0 && currentDirection == UP_DIRECTION) ||
		(newPosition.Y == len(m.MapPositions)-1 && currentDirection == DOWN_DIRECTION) {
		return currentPosition, false
	}
	return m.walk(newPosition, currentDirection)
}

func (m *SituationMap) isCollision(position Position) (ok bool) {
	return m.MapPositions[position.Y][position.X] == '#'
}

func (m *SituationMap) changeDirection(currentDirection Direction) Direction {
	switch currentDirection {
	case UP_DIRECTION:
		return RIGHT_DIRECTION
	case RIGHT_DIRECTION:
		return DOWN_DIRECTION
	case DOWN_DIRECTION:
		return LEFT_DIRECTION
	default:
		return UP_DIRECTION
	}
}
