package robot

import "testing"
import "fmt"

type MockSpeechModule struct {
	t    *testing.T
	said string
}

func (s MockSpeechModule) Say(text string) {
	fmt.Println(text)
	if s.said != text {
		s.t.Errorf("Expected %v but was %v", s.said, text)
	}
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
	r.SpeechModule = MockSpeechModule{t: t, said: "Hi. My name is EV3."}
	r.Greet()
}
