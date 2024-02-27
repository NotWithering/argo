package argo

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input        string
		expectedArgs []string
		expectedInc  bool
	}{
		{"\"hello world", []string{}, true},
		{"hello world", []string{"hello", "world"}, false},
		{"arg1 arg2 'arg 3'", []string{"arg1", "arg2", "arg 3"}, false},
		{"foo\\\\bar", []string{"foo\bar"}, false},
		{"unknown\\escape\\sequence", []string{"unknown\\escape\\sequence"}, false},
		{"new\\nline", []string{"new\nline"}, false},
		{"\"quoted string\"", []string{"quoted string"}, false},
		{"'single quoted'", []string{"single quoted"}, false},
		{"'single quoted' arg", []string{"single quoted", "arg"}, false},
		{"\"double quoted\" arg", []string{"double quoted", "arg"}, false},
		{"arg 'single quoted'", []string{"arg", "single quoted"}, false},
		{"arg \"double quoted\"", []string{"arg", "double quoted"}, false},
	}

	for _, tc := range testCases {
		arguments, incomplete := Parse(tc.input)
		if !reflect.DeepEqual(arguments, tc.expectedArgs) || incomplete != tc.expectedInc {
			t.Errorf("Parse(%q) = (%q, %t) expected (%q, %t)", tc.input, arguments, incomplete, tc.expectedArgs, tc.expectedInc)
		}
	}
}

func TestTerminates(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"'unterimated quote", false},
		{"\"unterminated double", false},
		{"\"terminated double\"", true},
		{"'terminated single'", true},
	}

	for _, tc := range testCases {
		terminates := Terminates(tc.input)
		if !reflect.DeepEqual(terminates, tc.expected) {
			t.Errorf("Terminates(%q) = %t, expected %t", tc.input, terminates, tc.expected)
		}
	}
}
