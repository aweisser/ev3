package robot

import (
	"fmt"
	"math"
	"strings"
)

// Specs is simple parameter object holding roboter specs. It should be used for better readability when constructing roboters
type Specs struct {
	WheelDiameter Centimeters
}

// Speaker enables speech capabilities
type Speaker interface {
	Speak(text string) error
}

// SoundPlayer enables sound playing capabilities
type SoundPlayer interface {
	Play(pathToAudio string) error
}

// Printer enables printing capabilities
type Printer interface {
	Print(text string) error
}

// Mover enables moving capabilities
type Mover interface {
	Move(distance Centimeters, tachometer Tachometer) error
}

// EventHandler enables handling for generic input events (like buttons)
type EventHandler interface {
	Handle(event interface{}) error
}

// Robot represents an executable EV3 device
type Robot struct {
	Name         string
	EnvMap       EnvironmentalMap
	Position     Position
	Tachometer   Tachometer
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
	err := r.tryMoveToPosition(steps)
	if err != nil {
		return err
	}
	distance := Centimeters(float64(steps) * float64(r.EnvMap.SquareSize))
	r.MoveModule.Move(distance, r.Tachometer)
	return nil
}

func (r *Robot) tryMoveToPosition(steps int) error {
	newPosition := r.Position
	absSteps := getAbsoluteSteps(steps)
	singleStep := getMovingDirection(steps)
	for i := 0; i < absSteps; i++ {
		newPosition = r.moveSingleStepFrom(newPosition, singleStep)
		if r.EnvMap.isObstacleAt(newPosition) {
			return fmt.Errorf("The robot can't move %v steps because there are obstacles in the way.Here's the current map: %v", steps, mapWithRobot(r))
		}
	}
	r.Position = newPosition
	return nil
}

func (r *Robot) moveSingleStepFrom(p Position, step int) Position {
	switch r.Position.Orientation {
	case North:
		p.Y = p.Y - step
	case East:
		p.X = p.X + step
	case South:
		p.Y = p.Y + step
	case West:
		p.X = p.X - step
	}
	return p
}

func getAbsoluteSteps(steps int) int {
	return int(math.Abs(float64(steps)))
}

func getMovingDirection(steps int) int {
	return steps / getAbsoluteSteps(steps)
}

// PrintEnvironment including the position of the robot
func (r *Robot) PrintEnvironment() {
	r.PrintModule.Print(mapWithRobot(r))
}

func mapWithRobot(r *Robot) string {
	rows := r.EnvMap.rows()
	offset := r.EnvMap.offset()
	rowIndexOfRobot := r.Position.Y + offset
	rows[rowIndexOfRobot] = placeRobotInRow(r, rows[rowIndexOfRobot])
	return strings.Join(rows, "\n")
}

func placeRobotInRow(r *Robot, row string) string {
	return fmt.Sprintf("%v%v%v", row[0:r.Position.X], r.Position.Orientation.String(), row[r.Position.X+1:len(row)])
}

// Handle arbitrary events
func (r *Robot) Handle(event interface{}) {
	r.EventHandler.Handle(event)
}
