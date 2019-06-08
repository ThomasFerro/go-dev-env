package standards

import (
	"log"
	"time"
)

// Debouncer A debouncer, managing a method's call
type Debouncer interface {
	Debounce(function func())
}

type defaultDebouncer struct {
	executeAfter int
	events       chan func()
}

func (d defaultDebouncer) init() {
	duration := time.Duration(d.executeAfter)
	var function func()
	t := time.NewTimer(duration)
	t.Stop()
	for {
		select {
		case newFunction := <-d.events:
			log.Println("Received a debouncer event with a new function")
			t.Reset(duration)
			function = newFunction
		case <-t.C:
			called := function != nil
			log.Printf("Debouncer function called ? %v\n", called)
			if called {
				function()
			}
		}
	}
}

func (d defaultDebouncer) Debounce(function func()) {
	d.events <- function
}

// NewDebouncer Create a new debouncer
func NewDebouncer(executeAfter int) Debouncer {
	events := make(chan func())
	debouncer := defaultDebouncer{
		executeAfter,
		events,
	}
	go debouncer.init()
	return debouncer
}
