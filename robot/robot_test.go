package robot

import "testing"
import "fmt"

type Mock struct {
	t        *testing.T
	expected interface{}
}

func (mock *Mock) verify(data interface{}) error {
	fmt.Println(data)
	if mock.expected != data {
		mock.t.Errorf("Expected %v but was %v", mock.expected, data)
	}
	return nil
}

func (mock *Mock) Speak(text string) error {
	return mock.verify(text)
}

func (mock *Mock) Print(text string) error {
	return mock.verify(text)
}

func (mock *Mock) Move(distance Centimeters, tachometer Tachometer) error {
	return mock.verify(distance)
}

func Test_TwoRobotsShouldHaveDifferentNames(t *testing.T) {
	r1 := Robot{}
	r1.Name = "Robot1"

	r2 := Robot{Name: "Robot2"}

	if r1.Name == r2.Name {
		t.Errorf("Robots should have different names. %v and %v", r1, r2)
	}
}

func Test_RobotShouldHaveGreetFunction(t *testing.T) {
	r := Robot{Name: "EV3"}
	r.SpeechModule = &Mock{t: t, expected: "Hi. My name is EV3."}
	r.Greet()
}

func Test_DirectionNorthShouldBePrintable(t *testing.T) {
	if North.String() != "▲" {
		t.Errorf("Expected %v but was %v", "▲", North)
	}
}

func Test_RobotShouldByAbleToLocateItselfOnAnEnvironmentalMap(t *testing.T) {
	r := Robot{}
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.Position = Position{X: 7, Y: 1, Orientation: North}

	expectedMap := `
###########
#      ▲  #
#         #
#         #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}

	r.PrintEnvironment()
}

func Test_RobotShouldDecreaseYPositionAfterMovePositiveStepsHeadingNorth(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(1)}
	r.Position = Position{X: 7, Y: 2, Orientation: North}
	r.Move(1)
	expectedPosition := Position{X: 7, Y: 1, Orientation: North}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldIncreasePositionYAfterMoveNegativeStepsHeadingNorth(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(-1)}
	r.Position = Position{X: 7, Y: 2, Orientation: North}
	r.Move(-1)
	expectedPosition := Position{X: 7, Y: 3, Orientation: North}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldIncreaseYPositionAfterMovePositiveStepsHeadingSouth(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(1)}
	r.Position = Position{X: 7, Y: 2, Orientation: South}
	r.Move(1)
	expectedPosition := Position{X: 7, Y: 3, Orientation: South}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldDecreasePositionYAfterMoveNegativeStepsHeadingSouth(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(-1)}
	r.Position = Position{X: 7, Y: 2, Orientation: South}
	r.Move(-1)
	expectedPosition := Position{X: 7, Y: 1, Orientation: South}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldIncreaseXPositionAfterMovePositiveStepsHeadingEast(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(1)}
	r.Position = Position{X: 3, Y: 2, Orientation: East}
	r.Move(1)
	expectedPosition := Position{X: 4, Y: 2, Orientation: East}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldDecreasePositionXAfterMoveNegativeStepsHeadingEast(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(-1)}
	r.Position = Position{X: 3, Y: 2, Orientation: East}
	r.Move(-1)
	expectedPosition := Position{X: 2, Y: 2, Orientation: East}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldDecreaseXPositionAfterMovePositiveStepsHeadingWest(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(1)}
	r.Position = Position{X: 3, Y: 2, Orientation: West}
	r.Move(1)
	expectedPosition := Position{X: 2, Y: 2, Orientation: West}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldIncreasePositionXAfterMoveNegativeStepsHeadingWest(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.MoveModule = &Mock{t: t, expected: Centimeters(-1)}
	r.Position = Position{X: 3, Y: 2, Orientation: West}
	r.Move(-1)
	expectedPosition := Position{X: 4, Y: 2, Orientation: West}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldMoveOnToTheNextFieldTowordsItsOrientation(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.Position = Position{X: 7, Y: 3, Orientation: North}

	expectedMap := `
###########
#         #
#         #
#      ▲  #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()

	r.MoveModule = &Mock{t: t, expected: Centimeters(2)}
	r.Move(2)

	expectedPosition := Position{X: 7, Y: 1, Orientation: North}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v but was %v", expectedMap, r.Position)
	}

	expectedMapAfterMove := `
###########
#      ▲  #
#         #
#         #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMapAfterMove}
	r.PrintEnvironment()
}

func Test_RobotCannotMoveThroughObstacles(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#         #
#         #
###########`
	r.Position = Position{X: 7, Y: 1, Orientation: North}

	expectedMap := `
###########
#      ▲  #
#         #
#         #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()

	r.MoveModule = &Mock{t: t, expected: Centimeters(2)}
	r.Move(1)

	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()
}

func Test_RobotCannotMoveThroughObstaclesInItsWayNorth(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#      #  #
#         #
###########`
	r.Position = Position{X: 7, Y: 3, Orientation: North}

	expectedMap := `
###########
#         #
#      #  #
#      ▲  #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()

	r.MoveModule = &Mock{t: t, expected: Centimeters(2)}
	err := r.Move(2)

	if err == nil {
		t.Errorf("Expected error message")
	}

	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()
}

func Test_RobotCannotMoveThroughObstaclesInItsWayBack(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#      #  #
#         #
###########`
	r.Position = Position{X: 7, Y: 1, Orientation: North}

	expectedMap := `
###########
#      ▲  #
#      #  #
#         #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()

	r.MoveModule = &Mock{t: t, expected: Centimeters(2)}
	err := r.Move(-2)

	if err == nil {
		t.Errorf("Expected error message")
	}

	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()
}

func Test_RobotCannotMoveThroughObstaclesInItsWayEast(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	r.EnvMap.Map = `
###########
#         #
#      #  #
#         #
###########`
	r.Position = Position{X: 6, Y: 2, Orientation: East}

	expectedMap := `
###########
#         #
#     ▶#  #
#         #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()

	r.MoveModule = &Mock{t: t, expected: Centimeters(2)}
	err := r.Move(2)

	if err == nil {
		t.Errorf("Expected error message")
	}

	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()
}

func Test_RobotWithoutEnvironmentalMapCanNotMoveAtAll(t *testing.T) {
	r := Robot{}
	r.EnvMap.SquareSize = 1
	err := r.Move(1)
	if err == nil {
		t.Errorf("Move should return error on missing env map")
	}
}
