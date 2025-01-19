package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input  string
		output []string
	}
	tests := []testCase{
		{"hello world", []string{"hello", "world"}},
		{"do something quickly", []string{"do", "something", "quickly"}},
		{"do SOMETHING SLOWLY", []string{"do", "something", "slowly"}},
		{"  do aNoThEr Thing     SLOWLY", []string{"do", "another", "thing", "slowly"}},
	}

	for _, test := range tests {
		result := cleanInput(test.input)
		if reflect.DeepEqual(result, test.output) == false {
			t.Errorf("%v != %v", result, test.output)
		}
	}

}
