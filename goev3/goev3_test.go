package goev3

import "testing"

func Test_speedShouldBe0ForZeroSteps(t *testing.T) {
	var expectedSpeed int16
	if speed(0) != expectedSpeed {
		t.Errorf("Speed should be %v", expectedSpeed)
	}
}

func Test_speedShouldBe100ForPositiveSteps(t *testing.T) {
	var expectedSpeed int16 = 100
	if speed(1) != expectedSpeed {
		t.Errorf("Speed should be %v", expectedSpeed)
	}
}

func Test_speedShouldBeMinus100ForNegativeSteps(t *testing.T) {
	var expectedSpeed int16 = -100
	if speed(-1) != expectedSpeed {
		t.Errorf("Speed should be %v", expectedSpeed)
	}
}
