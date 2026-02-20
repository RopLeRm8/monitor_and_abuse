package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"git.sr.ht/~jackmordaunt/go-toast"
	"github.com/mitchellh/go-ps"
)

var MAX_TIME = 60 * 120 // 120 Minutes = 2h
var DELAY = 60 * 10

type Process string

const (
	PROCESS_CHROME Process = "overwatch.exe"
)

var BLOCKED_PROCESSES = map[Process]string{
	"overwatch.exe": "Overwatch",
}

var TIMERS = map[Process]time.Time{
	PROCESS_CHROME: time.Now(),
}

//go:embed check.png
var icon []byte

func notify(process Process) {
	appName := BLOCKED_PROCESSES[process]
	notification := toast.Notification{
		AppID: "Monitor And Abuse Control System",
		Title: fmt.Sprintf("Too much time on %s", appName),
		Body:  fmt.Sprintf("You spent more than 2 hours on %s, get to work!", appName),
		Icon:  "check.png",
		Loop:  false,
	}
	notification.Push()
}

func main() {

	for {
		time.Sleep(time.Second * time.Duration(DELAY))

		procs, err := ps.Processes()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		for _, proc := range procs {
			exe := Process(strings.ToLower(proc.Executable()))

			if _, timerExists := TIMERS[exe]; !timerExists {
				continue
			}

			if time.Now().Before(TIMERS[exe]) {
				continue
			}

			notify(exe)
			TIMERS[exe] = TIMERS[exe].Add(time.Second * time.Duration(MAX_TIME))

		}

	}

}
