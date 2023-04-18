package main

import (
	"os"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	// Set up test cases
	testCases := []struct {
		key   string
		value string
		err   error
	}{
		{key: "FOO", value: "bar", err: nil},
		{key: "coolVar", value: "Oh Yes", err: nil},
		{key: "SECRET_BUDDY", value: "{\"current\":{\"KEY1\":\"VALUE4\",\"KEY2\":\"VALUE4\"},\"previous\":{\"KEY1\":\"VALUE3\",\"KEY2\":\"VALUE3\"}}", err: nil},
	}

	// Set up environment variables for test cases
	os.Setenv("FOO", "bar")
	os.Setenv("coolVar", "Oh Yes")
	os.Setenv("SECRET_BUDDY", "{\"current\":{\"KEY1\":\"VALUE4\",\"KEY2\":\"VALUE4\"},\"previous\":{\"KEY1\":\"VALUE3\",\"KEY2\":\"VALUE3\"}}")

	// Clean up environment variables after tests finish
	defer os.Unsetenv("FOO")
	defer os.Unsetenv("coolVar")
	defer os.Unsetenv("SECRET_BUDDY")

	// Run tests
	for _, tc := range testCases {
		actual, err := GetEnvVar(tc.key)

		// Check error values
		if err != tc.err {
			t.Errorf("Unexpected error. Expected: %v, got: %v", tc.err, err)
		}

		// Check return values
		if actual != tc.value {
			t.Errorf("Unexpected result. Expected: %v, got: %v", tc.value, actual)
		}
	}
}

func TestExportEnvVarsFromMap(t *testing.T) {
	// Set up test case
	input := `{"current": {"FOO": "bar", "KEY": "VALUE4"}, "previous": {"FOO": "noooo", "KEY": "VALUE3"}}`
	expectedOutput := map[string]string{"FOO": "bar", "KEY": "VALUE4"}

	// Run function
	actualOutput, err := ExportEnvVarsFromMap(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check output
	for key, expectedValue := range expectedOutput {
		actualValue, ok := actualOutput[key]
		if !ok {
			t.Errorf("Missing key in output: %q", key)
		} else if actualValue != expectedValue {
			t.Errorf("Unexpected value for key %q. Expected: %q, got: %q", key, expectedValue, actualValue)
		}
	}
}
