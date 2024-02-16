package main

import "testing"

// Test to see if name matches want "Susan"
func Test_sayHello(t *testing.T) {
	name := "Susan"
	want := "Hello Susan"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
