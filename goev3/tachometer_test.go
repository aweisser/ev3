package goev3

import "testing"

func Test_tachometerForOneCentimeterDistance(t *testing.T) {
	tachometer := tachometer{distance: 1.0, wheelDiameter: 3.5}
	degrees := tachometer.counts()
	expectedDegrees := 32.74044543604704
	if degrees != expectedDegrees {
		t.Errorf("Tachometer counts of %v should be %v but was %v", tachometer, expectedDegrees, degrees)
	}
}
