package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHello(t *testing.T) {
	name := "Andrew"
	want := "Hi, Andrew. Welcome!"
	msg := Hello(name)
	assert.Equal(t, msg, want)
}
