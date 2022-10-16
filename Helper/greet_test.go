package helper

import (
	"testing"
)

func TestGreet(t *testing.T) {
	greeting, greetingLength := greet("Alice")

	Equal(t, "Hello Alice", greeting)
	Equal(t, 12, greetingLength)
}
