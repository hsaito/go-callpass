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
