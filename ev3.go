package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mattrajca/GoEV3/TTS"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is EV3 mission control. Start a mission!")
}

func exit(w http.ResponseWriter, r *http.Request) {
	defer os.Exit(0)
	io.WriteString(w, "Stopping EV3 mission control. CU soon!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	TTS.Speak("This is EV3. Ready for action.")
}

func main() {
	fmt.Printf("Starting EV3 command control ...")
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/exit", exit)
	http.ListenAndServe(":9000", nil)
}
