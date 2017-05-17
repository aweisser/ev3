package robot

import "testing"

func Test_tachometerForOneCentimeterDistance(t *testing.T) {
	tachometer := Tachometer{WheelDiameter: 3.5, CountPerRot: 360}
	degrees := tachometer.CountsForDistance(1.0)
	expectedDegrees := 32.74044543604704
	if degrees != expectedDegrees {
		t.Errorf("Tachometer counts of %v should be %v but was %v", tachometer, expectedDegrees, degrees)
	}
}
