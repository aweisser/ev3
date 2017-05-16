package robot

import "testing"
import "fmt"

type Mock struct {
	t        *testing.T
	expected string
}

func (mock *Mock) verify(text string) {
	fmt.Println(text)
	if mock.expected != text {
		mock.t.Errorf("Expected %v but was %v", mock.expected, text)
	}
}

func (mock *Mock) Say(text string) {
	mock.verify(text)
}

func (mock *Mock) Print(text string) {
	mock.verify(text)
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
	r.Position = Position{X: 6, Y: 3, Orientation: North}

	expectedMap := `
###########
#         #
#         #
#      ▲  #
###########`
	r.PrintModule = &Mock{t: t, expected: expectedMap}

	r.PrintEnvironment()
}
