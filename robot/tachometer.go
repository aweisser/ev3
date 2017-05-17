package robot

import (
	"math"
)

// Tachometer for a given wheel size and tachometer pulse counts per one rotation
// Unlike other MINDSTORMS software, the units of measurement used are in tachometer counts rather than rotations or degrees.
// For the NXT and EV3 motors, one pulse of the tachometer = one degree. So the tachometer pulse count per one rotation is 360.
type Tachometer struct {
	WheelDiameter Centimeters
	CountPerRot   float64 // tachometer pulse counts per one rotation
}

// CountsForDistance calculates the so called Tachometer counts for a given distence based on the wheel size of this Tachometer instance.
// Math taken from http://ev3lessons.com/resources/wheelconverter/
func (t *Tachometer) CountsForDistance(distance Centimeters) float64 {
	wheelCircumference := math.Pi * float64(t.WheelDiameter)
	numberOfWheelRotations := float64(distance) / wheelCircumference
	numberOfDegrees := numberOfWheelRotations * t.CountPerRot
	return numberOfDegrees
}
