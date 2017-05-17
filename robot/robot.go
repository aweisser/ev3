package robot

import (
	"fmt"
	"strings"
)

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

// Speaker enables speech capabilities
type Speaker interface {
	Say(text string)
}

// Printer enables printing capabilities
type Printer interface {
	Print(text string)
}

// Mover enables moving capabilities
type Mover interface {
	Move(steps int)
}

// Robot represents an executable EV3 device
type Robot struct {
	Name         string
	EnvMap       string
	Position     Position
	SpeechModule Speaker
	PrintModule  Printer
	MoveModule   Mover
}

// Greet with your name
func (r *Robot) Greet() {
	r.SpeechModule.Say(fmt.Sprintf("Hi. My name is %v.", r.Name))
}

// Move forward (if steps are positive) or backward (if steps are negative)
func (r *Robot) Move(steps int) {
	r.MoveModule.Move(steps)
	switch r.Position.Orientation {
	case North:
		r.Position.Y -= steps
	case East:
		r.Position.X += steps
	case South:
		r.Position.Y += steps
	case West:
		r.Position.X -= steps
	}
}

// PrintEnvironment prints the current environment and the position of the robot
func (r *Robot) PrintEnvironment() {
	rows := strings.Split(r.EnvMap, "\n")

	// ignore first leading blank line (only for better experience with multiline strings)
	offset := 0
	if rows[0] == "" {
		offset = 1
	}
	rowIndexOfRobot := r.Position.Y + offset
	rows[rowIndexOfRobot] = placeRobot(r, rows[rowIndexOfRobot])
	r.PrintModule.Print(strings.Join(rows, "\n"))
}

func placeRobot(r *Robot, row string) string {
	return fmt.Sprintf("%v%v%v", row[0:r.Position.X+1], r.Position.Orientation.String(), row[r.Position.X+2:len(row)])
}
