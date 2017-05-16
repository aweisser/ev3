package robot

import (
	"fmt"
)

// SpeechModule enables speech abilities
type SpeechModule interface {
	Say(text string)
}

// Robot represents an executable EV3 device
type Robot struct {
	Name         string
	SpeechModule SpeechModule
}

// Greet with your name
func (r *Robot) Greet() {
	r.SpeechModule.Say(fmt.Sprintf("Hi. My name is %v.", r.Name))
}
