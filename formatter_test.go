package golden

import (
	"io/ioutil"
	"testing"
)

func TestFormatString(t *testing.T) {
	input := "Hello, World!"
	expectedOutput, err := ioutil.ReadFile("expected_output.golden")
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	actualOutput := FormatString(input)

	if string(actualOutput) != string(expectedOutput) {
		t.Errorf("Unexpected output:\nExpected: %s\nActual: %s", expectedOutput, actualOutput)
	}
}
