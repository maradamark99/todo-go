package test

import "testing"

func TestSmth(t *testing.T) {
	if 1 == 2 {
		t.Errorf("1 is equal to 2 for some reason.")
	}
}
