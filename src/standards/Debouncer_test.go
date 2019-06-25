package standards_test

import (
	"testing"
	"time"

	"github.com/go-dev-env/standards"
)

func TestShouldCallTheFunctionAfterTheProvidedTime(t *testing.T) {
	called := false
	debouncedMethod := standards.Debounce(func() {
		called = true
	}, 1)

	debouncedMethod()

	time.Sleep(10 * time.Millisecond)

	if called == false {
		t.Error("The method was not called")
	}
}

func TestShouldNotCallTheFunctionBeforeTheProvidedTime(t *testing.T) {
	called := false
	debouncedMethod := standards.Debounce(func() {
		called = true
	}, 10)

	debouncedMethod()

	time.Sleep(1 * time.Millisecond)

	if called == true {
		t.Error("The method was called too early")
	}
}

func TestShouldCallTheFunctionOnlyOnceAfterTheProvidedTime(t *testing.T) {
	executionCount := 0
	debouncedMethod := standards.Debounce(func() {
		executionCount++
	}, 1)

	debouncedMethod()
	debouncedMethod()
	debouncedMethod()
	debouncedMethod()

	time.Sleep(5 * time.Millisecond)

	if executionCount != 1 {
		t.Errorf("The method was not called only once, called %v time(s)", executionCount)
	}
}

func TestShouldBeAbleToCallTheFunctionAgainAfterTheTimer(t *testing.T) {
	executionCount := 0
	debouncedMethod := standards.Debounce(func() {
		executionCount++
	}, 5)

	debouncedMethod()

	time.Sleep(10 * time.Millisecond)

	debouncedMethod()

	time.Sleep(10 * time.Millisecond)

	if executionCount != 2 {
		t.Errorf("The method was not called twice, called %v time(s)", executionCount)
	}
}
