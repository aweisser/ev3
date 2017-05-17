package goev3

import (
	"math"

	"github.com/aweisser/ev3/robot"
)

type tachometer struct {
	distance      robot.Centimeters
	wheelDiameter robot.Centimeters
}

// taken from http://ev3lessons.com/resources/wheelconverter/
func (t *tachometer) counts() float64 {
	wheelCircumference := math.Pi * t.wheelDiameter
	numberOfWheelRotations := t.distance / wheelCircumference
	numberOfDegrees := numberOfWheelRotations * 360
	return float64(numberOfDegrees)
}
