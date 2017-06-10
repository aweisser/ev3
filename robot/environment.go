package robot

import "strings"

type Centimeters float64

// EnvironmentalMap in which the robot acts
type EnvironmentalMap struct {
	Map        string
	SquareSize Centimeters
}

func (e *EnvironmentalMap) rows() []string {
	return strings.Split(e.Map, "\n")
}

func (e *EnvironmentalMap) offset() int {
	// ignore first leading blank line (only for better experience with multiline strings)
	offset := 0
	if e.rows()[0] == "" {
		offset = 1
	}
	return offset
}

func (e *EnvironmentalMap) isObstacleAt(p Position) bool {
	return e.isOutsideMap(p) || e.isObstacleInMap(p)
}

func (e *EnvironmentalMap) isObstacleInMap(p Position) bool {
	return e.rows()[p.Y+e.offset()][p.X] == '#'
}

func (e *EnvironmentalMap) isOutsideMap(p Position) bool {
	rows := e.rows()
	offset := e.offset()
	return isOutsideVertical(p, rows, offset) || isOutsideHorizontal(p, rows, offset)
}

func isOutsideVertical(p Position, rows []string, offset int) bool {
	maxRowIndex := len(rows) - offset - 1
	return p.Y < 0 || p.Y > maxRowIndex
}

func isOutsideHorizontal(p Position, rows []string, offset int) bool {
	row := rows[p.Y+offset]
	maxColumnIndexInRow := len(row) - 1
	return p.X < 0 || p.X > maxColumnIndexInRow
}

// The Direction of the robot
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "▲"
	case East:
		return "▶"
	case South:
		return "▼"
	case West:
		return "◀"
	default:
		return "X"
	}
}

// Position of the robot in the environmental map including Orientation
type Position struct {
	X           int
	Y           int
	Orientation Direction
}
