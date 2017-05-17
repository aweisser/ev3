package robot

import "testing"
import "fmt"

type Mock struct {
	t        *testing.T
	expected interface{}
}

func (mock *Mock) verify(data interface{}) {
	fmt.Println(data)
	if mock.expected != data {
		mock.t.Errorf("Expected %v but was %v", mock.expected, data)
	}
}

func (mock *Mock) Say(text string) {
	mock.verify(text)
}

func (mock *Mock) Print(text string) {
	mock.verify(text)
}

func (mock *Mock) Move(steps int) {
	mock.verify(steps)
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
	r.EnvMap = `
###########
#         #
#         #
#         #
###########`
	r.Position = Position{X: 6, Y: 1, Orientation: North}

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
	r.MoveModule = &Mock{t: t, expected: 1}
	r.Position = Position{X: 6, Y: 2, Orientation: North}
	r.Move(1)
	expectedPosition := Position{X: 6, Y: 1, Orientation: North}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldIncreasePositionYAfterMoveNegativeStepsHeadingNorth(t *testing.T) {
	r := Robot{}
	r.MoveModule = &Mock{t: t, expected: -1}
	r.Position = Position{X: 6, Y: 2, Orientation: North}
	r.Move(-1)
	expectedPosition := Position{X: 6, Y: 3, Orientation: North}
	if r.Position != expectedPosition {
		t.Errorf("Expected %v, but was %v", expectedPosition, r.Position)
	}
}

func Test_RobotShouldMoveOnToTheNextFieldTowordsItsOrientation(t *testing.T) {
	r := Robot{}
	r.EnvMap = `
###########
#         #
#         #
#         #
###########`
	r.Position = Position{X: 6, Y: 3, Orientation: North}

	expectedMap := `
###########
#         #
#         #
#      ▲  #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}
	r.PrintEnvironment()

	r.MoveModule = &Mock{t: t, expected: 2}
	r.Move(2)

	expectedPosition := Position{X: 6, Y: 1, Orientation: North}
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
