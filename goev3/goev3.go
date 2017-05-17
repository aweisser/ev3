package goev3

import (
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

func (e *goEV3Engine) Speak(text string) {
	TTS.Speak(text)
}

func (e *goEV3Engine) Handle(event interface{}) {
	switch event {
	case WAIT_FOR_ESCAPE_BUTTON:
		Button.Wait(Button.Escape)
	}
}

func (e *goEV3Engine) Move(steps int) {

	// Unlike other MINDSTORMS software, the units of measurement used are in tachometer counts rather than rotations or degrees.
	// For the NXT and EV3 motors, one pulse of the tachometer = one degree.

	// check current position
	//positionLeftWheel := Motor.CurrentPosition(left_wheel)
	//positionRightWheel := Motor.CurrentPosition(right_wheel)

	// calc final position
	//positionLeftWheel += steps * 360

	// run wheels until final position has been reached
	speed := speed(steps)
	Motor.Run(left_wheel, speed)
	Motor.Run(right_wheel, speed)
}

func speed(steps int) int16 {
	var speed int16
	switch {
	case steps == 0:
		steps = 0
	case steps < 0:
		speed = -100
	case steps > 0:
		speed = 100
	}
	return speed
}
