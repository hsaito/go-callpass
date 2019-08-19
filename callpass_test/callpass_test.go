package callpass_test

import (
	"go-callpass/callpass"
	"testing"
)

func TestCallpass(t *testing.T) {
	result := callpass.GetCallpass("TESTING")

	if result != 31421 {
		t.Errorf("Wrong callpass result.")
	}
}

func TestCallpassValidation(t *testing.T) {
	result := callpass.CheckCallpass("TESTING", 31421)

	if !result {
		t.Errorf("Validation failed.")
	}
}
