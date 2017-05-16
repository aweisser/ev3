package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aweisser/GoEV3/Button"
	"github.com/aweisser/GoEV3/TTS"
)

func main() {
	defer Button.Wait(Button.Escape)
	fmt.Printf("Starting EV3 restless command control ...")
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is EV3 mission control. Start a mission!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	TTS.Speak("This is EV3. Ready for action.")
}
