package goev3

import (
	"github.com/aweisser/GoEV3/Button"
	"github.com/aweisser/GoEV3/TTS"
	"github.com/aweisser/ev3/robot"
)

// Create a robot with goev3 engine
func Create() robot.Robot {
	return robot.Robot{
		Name:         "GoEV3",
		SpeechModule: new(goEV3Engine),
	}
}

type GoEV3Event int

const (
	WAIT_FOR_ESCAPE_BUTTON GoEV3Event = iota
)

type goEV3Engine struct {
}

func (e *goEV3Engine) Say(text string) {
	TTS.Speak(text)
}

func (e *goEV3Engine) Handle(event interface{}) {
	switch event {
	case WAIT_FOR_ESCAPE_BUTTON:
		Button.Wait(Button.Escape)
	}

}
