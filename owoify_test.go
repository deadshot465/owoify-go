package owoify_go

import "testing"

func TestOwoify(t *testing.T) {
	if got := Owoify("This is the string to owo! Kinda cute, isn't it?", "owo"); len(got) == 0 {
		t.Error("Failed to owoify the string.")
	}
}
