package callpass_test

import (
	"go-callpass/callpass"
	"testing"
)

func TestCallPass(t *testing.T) {
	result := callpass.GetCallPass("TESTING")

	if result != 31421 {
		t.Errorf("Wrong callpass result.")
	}
}

func TestCallPassValidation(t *testing.T) {
	result := callpass.CheckCallPass("TESTING", 31421)

	if !result {
		t.Errorf("Validation failed.")
	}
}
