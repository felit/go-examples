package gotest

import "testing"

func TestBasic(test *testing.T) {
	grade := "D"
	if grade != "D" {
		test.Error("Test Case failed.")
	}

}
