package standards

import (
	"log"
	"time"
)

func Debounce(function func(), executeAfter int) func() {
	duration := time.Duration(executeAfter) * time.Millisecond
	t := time.NewTimer(duration)
	t.Stop()

	go (func() {
		for {
			select {
			case <-t.C:
				log.Println("Executing the debounced method")
				go function()
			}
		}
	})()

	return func() {
		log.Println("Reset the debouncer timer")
		t.Reset(duration)
	}
}
