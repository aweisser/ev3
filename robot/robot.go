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

// Speaker enables speech abilities
type Speaker interface {
	Say(text string)
}

// Printer enables printing abilities
type Printer interface {
	Print(text string)
}

// Robot represents an executable EV3 device
type Robot struct {
	Name         string
	EnvMap       string
	Position     Position
	SpeechModule Speaker
	PrintModule  Printer
}

// Greet with your name
func (r *Robot) Greet() {
	r.SpeechModule.Say(fmt.Sprintf("Hi. My name is %v.", r.Name))
}

// PrintEnvironment prints the current environment and the position of the robot
func (r *Robot) PrintEnvironment() {
	lines := strings.Split(r.EnvMap, "\n")
	offset := 0
	if lines[0] == "" {
		offset = 1
	}
	line := lines[r.Position.Y]
	newLine := fmt.Sprintf("%v%v%v", line[0:r.Position.X+1], r.Position.Orientation.String(), line[r.Position.X+2:len(line)])
	lines[r.Position.Y+offset] = newLine
	r.PrintModule.Print(strings.Join(lines, "\n"))
}
