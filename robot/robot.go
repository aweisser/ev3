package robot

import (
	"fmt"
	"strings"
)

// Speaker enables speech capabilities
type Speaker interface {
	Speak(text string)
}

// SoundPlayer enables sound playing capabilities
type SoundPlayer interface {
	Play(pathToAudio string)
}

// Printer enables printing capabilities
type Printer interface {
	Print(text string)
}

// Mover enables moving capabilities
type Mover interface {
	Move(steps int)
}

// EventHandler enables handling for generic input events (like buttons)
type EventHandler interface {
	Handle(event interface{})
}

// Robot represents an executable EV3 device
type Robot struct {
	Name         string
	EnvMap       EnvironmentalMap
	Position     Position
	SpeechModule Speaker
	PrintModule  Printer
	MoveModule   Mover
	EventHandler EventHandler
}

// Greet with your name
func (r *Robot) Greet() {
	r.SpeechModule.Speak(fmt.Sprintf("Hi. My name is %v.", r.Name))
}

// Move forward (if steps are positive) or backward (if steps are negative)
func (r *Robot) Move(steps int) error {
	if r.EnvMap.Map == "" {
		return fmt.Errorf("The robot can't move because no environmental map is given")
	}
	newPosition := r.Position
	switch r.Position.Orientation {
	case North:
		newPosition.Y = r.Position.Y - steps
	case East:
		newPosition.X = r.Position.X + steps
	case South:
		newPosition.Y = r.Position.Y + steps
	case West:
		newPosition.X = r.Position.X - steps
	}
	if !r.EnvMap.isObstacle(newPosition) {
		r.MoveModule.Move(steps)
		r.Position = newPosition
	}
	return nil
}

// PrintEnvironment including the position of the robot
func (r *Robot) PrintEnvironment() {
	rows := r.EnvMap.rows()
	offset := r.EnvMap.offset()
	rowIndexOfRobot := r.Position.Y + offset
	rows[rowIndexOfRobot] = placeRobotInRow(r, rows[rowIndexOfRobot])
	r.PrintModule.Print(strings.Join(rows, "\n"))
}

func placeRobotInRow(r *Robot, row string) string {
	return fmt.Sprintf("%v%v%v", row[0:r.Position.X+1], r.Position.Orientation.String(), row[r.Position.X+2:len(row)])
}

// Handle arbitrary events
func (r *Robot) Handle(event interface{}) {
	r.EventHandler.Handle(event)
}
