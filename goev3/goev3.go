package goev3

import (
	"fmt"

	"github.com/aweisser/GoEV3/Button"
	"github.com/aweisser/GoEV3/Motor"
	"github.com/aweisser/GoEV3/TTS"
	"github.com/aweisser/ev3/robot"
)

// Create a robot with goev3 engine
func Create() robot.Robot {
	goEV3EngineInstance := new(goEV3Engine)
	return robot.Robot{
		Name:         "GoEV3",
		SpeechModule: goEV3EngineInstance,
		MoveModule:   goEV3EngineInstance,
	}
}

type eventType int

const (
	WAIT_FOR_ESCAPE_BUTTON eventType = iota
)

const (
	left_wheel  Motor.OutPort = Motor.OutPortA
	right_wheel               = Motor.OutPortB
)

type goEV3Engine struct {
}

func (e *goEV3Engine) Speak(text string) error {
	TTS.Speak(text)
	return nil
}

func (e *goEV3Engine) Handle(event interface{}) error {
	switch event {
	case WAIT_FOR_ESCAPE_BUTTON:
		Button.Wait(Button.Escape)
	default:
		return fmt.Errorf("Unknown event %v", event)
	}
	return nil
}

func (e *goEV3Engine) Move(distance robot.Centimeters) error {

	if distance == 0 {
		return nil
	}

	if distance < 0 {
		return fmt.Errorf("Moving back is not implemented yet")
	}

	// check current position
	positionLeftWheel := Motor.CurrentPosition(left_wheel)
	positionRightWheel := Motor.CurrentPosition(right_wheel)

	// calc final position
	tm := robot.Tachometer{WheelDiameter: 3.2, PulsesPerDegree: 1.0} // TODO hardcoded
	degreeToMove := int32(tm.CountsForDistance(distance))
	positionLeftWheel += degreeToMove
	positionRightWheel += degreeToMove

	// run wheels ...
	Motor.Run(left_wheel, 100)
	Motor.Run(right_wheel, 100)

	// ... until final position has been reached
	for {
		if Motor.CurrentPosition(left_wheel) >= positionLeftWheel && Motor.CurrentPosition(right_wheel) >= positionRightWheel {
			break
		}
	}
	Motor.Stop(left_wheel)
	Motor.Stop(right_wheel)
	return nil
}

func speed(distance robot.Centimeters) int16 {
	var speed int16
	switch {
	case distance == 0:
		speed = 0
	case distance < 0:
		speed = -100
	case distance > 0:
		speed = 100
	}
	return speed
}
