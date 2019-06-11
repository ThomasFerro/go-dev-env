package standards

import (
	"log"
	"time"
)

func Debounce(function func(), executeAfter int) func() {
	debounceTriggered := make(chan bool)

	log.Println("Initiate the debounced method")
	duration := time.Duration(executeAfter) * time.Millisecond
	t := time.NewTimer(duration)
	t.Stop()

	startDebounceLoop := func() {
		for {
			select {
			case <-debounceTriggered:
				log.Println("Received a debouncer event, reset the timer")
				t.Reset(duration)
			case <-t.C:
				log.Println("Executing the debounced method")
				go function()
			}
		}
	}

	go startDebounceLoop()

	return func() {
		log.Println("RETURNED FUNCTION CALLED")
		debounceTriggered <- true
	}
}
