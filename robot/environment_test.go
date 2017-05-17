package robot

import "testing"

func Test_PositionYLesserZeroShouldBeOutsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 0, Y: -1}
	if !m.isOutsideMap(p) {
		t.Errorf("%v should be outside map %v", p, m)
	}
}

func Test_PositionYEqualsZeroShouldBeInsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 0, Y: 0}
	if m.isOutsideMap(p) {
		t.Errorf("%v should be inside map %v", p, m)
	}
}

func Test_PositionYGreaterThanMaxRowIndexShouldBeOutsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 0, Y: 3}
	if !m.isOutsideMap(p) {
		t.Errorf("%v should be outside map %v", p, m)
	}
}

func Test_PositionYGreaterThanMaxRowIndexWithoutOffsetRowShouldBeOutsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `#########
#       #
#########`
	p := Position{X: 0, Y: 3}
	if !m.isOutsideMap(p) {
		t.Errorf("%v should be outside map %v", p, m)
	}
}

func Test_PositionXLesserZeroShouldBeOutsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: -1, Y: 0}
	if !m.isOutsideMap(p) {
		t.Errorf("%v should be outside map %v", p, m)
	}
}

func Test_PositionXEqualsZeroShouldBeInsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 0, Y: 0}
	if m.isOutsideMap(p) {
		t.Errorf("%v should be inside map %v", p, m)
	}
}

func Test_PositionGreaterThanMaxColumnIndexInRowShouldBeOutsideMap(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 9, Y: 1}
	if !m.isOutsideMap(p) {
		t.Errorf("%v should be outside map %v", p, m)
	}
}

func Test_PositionOutsideMapHasAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 9, Y: 1}
	if !m.isObstacle(p) {
		t.Errorf("%v should be an obstacle in map %v", p, m)
	}
}

func Test_UpperWallShouldBeAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 4, Y: 0}
	if !m.isObstacle(p) {
		t.Errorf("%v should be an obstacle in map %v", p, m)
	}
}

func Test_LeftWallShouldBeAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 0, Y: 1}
	if !m.isObstacle(p) {
		t.Errorf("%v should be an obstacle in map %v", p, m)
	}
}

func Test_RightWallShouldBeAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 8, Y: 1}
	if !m.isObstacle(p) {
		t.Errorf("%v should be an obstacle in map %v", p, m)
	}
}

func Test_BottomWallShouldBeAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 4, Y: 2}
	if !m.isObstacle(p) {
		t.Errorf("%v should be an obstacle in map %v", p, m)
	}
}

func Test_FreeObstacleInTheMiddleOfTheMapShouldBeAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#   #   #
#########`
	p := Position{X: 4, Y: 1}
	if !m.isObstacle(p) {
		t.Errorf("%v should be an obstacle in map %v", p, m)
	}
}

func Test_FreeSpaceInTheMiddleOfTheMapShouldNotBeAnObstacle(t *testing.T) {
	m := EnvironmentalMap{}
	m.Map = `
#########
#       #
#########`
	p := Position{X: 4, Y: 1}
	if m.isObstacle(p) {
		t.Errorf("%v should not be an obstacle in map %v", p, m)
	}
}
