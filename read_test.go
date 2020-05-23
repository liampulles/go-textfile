package textfile

import (
	"fmt"
	"path"
	"reflect"
	"testing"
)

func TestRead_ReadAllLines_GivenExistingFile_ShouldReadLines(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		filenameFixture string
		expected        []string
	}{
		// Blank file
		{
			"blank.txt",
			nil,
		},
		// Blank lines
		{
			"blanklines.txt",
			lines("", ""),
		},
		// Normal case
		{
			"normal.txt",
			lines("here lies", "some", "text!"),
		},
		// Normal case with additional newline
		{
			"normalnewline.txt",
			lines("you may", "find", "some text", "here"),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s -> %v",
			test.filenameFixture, test.expected), func(t *testing.T) {
			// Setup fixture
			resolvedPath := path.Join("testdata", test.filenameFixture)

			// Exercise SUT
			actual, err := ReadAllLines(resolvedPath)

			// Verify result
			if err != nil {
				t.Errorf("Encountered error\n%#v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Unexpected Result.\nActual: %#v\nExpected: %#v", actual, test.expected)
			}
		})
	}
}

func TestRead_ReadAllLines_FailingCases(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		filenameFixture string
		expectedErr     error
	}{
		// Non-existent file
		{
			"non existent",
			fmt.Errorf("open %s: no such file or directory", path.Join("testdata", "non existent")),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s -> %v",
			test.filenameFixture, test.expectedErr), func(t *testing.T) {
			// Setup fixture
			resolvedPath := path.Join("testdata", test.filenameFixture)

			// Exercise SUT
			actual, err := ReadAllLines(resolvedPath)

			// Verify result
			if err == nil {
				t.Errorf("Expected error, but none returned")
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("Unexpected Result.\nActual: %#v\nExpected: %#v", err.Error(), test.expectedErr.Error())
			}
			if !reflect.DeepEqual(actual, []string(nil)) {
				t.Errorf("Unexpected Result.\nActual: %#v\nExpected: %#v", actual, []string(nil))
			}
		})
	}
}

func lines(l ...string) []string {
	return l
}
