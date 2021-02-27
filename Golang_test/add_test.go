package main

import "testing"

func TestAdd(t *testing.T) {
	var testcvar = add(1, 2)

	if testcvar != 3 {
		t.Errorf("the add function failed, expected 3 and got %v", testcvar)
	}
}
