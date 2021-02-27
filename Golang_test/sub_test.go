package main

import "testing"

func TestSub(t *testing.T) {
	var testVar = sub(3, 2)

	if testVar != 1 {
		t.Errorf("the subtraction function failed, expected 1 and got %v", testVar)
	}
}
