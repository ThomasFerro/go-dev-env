package standards

import (
	"testing"
	"time"
)

func TestShouldCallTheFunctionAfterTheProvidedTime(t *testing.T) {
	called := false
	debouncedMethod := Debounce(func() {
		called = true
	}, 10)

	debouncedMethod()

	time.Sleep(100 * time.Millisecond)

	if called == false {
		t.Error("The method was not called")
	}
}

func TestShouldNotCallTheFunctionBeforeTheProvidedTime(t *testing.T) {
	called := false
	debouncedMethod := Debounce(func() {
		called = true
	}, 500)

	debouncedMethod()

	time.Sleep(1 * time.Millisecond)

	if called == true {
		t.Error("The method was called too early")
	}
}

func TestShouldCallTheFunctionOnlyOnceAfterTheProvidedTime(t *testing.T) {
	callCount := 0
	debouncedMethod := Debounce(func() {
		callCount++
	}, 1)

	debouncedMethod()
	debouncedMethod()
	debouncedMethod()
	debouncedMethod()

	time.Sleep(5 * time.Millisecond)

	if callCount != 1 {
		t.Errorf("The method was not called only once, called %v time(s)", callCount)
	}
}
