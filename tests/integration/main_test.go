package tests

import (
	"cli-json-formatter/internal/service"
	"fmt"
	"reflect"
	"testing"
)

func TestFormatStringToJson(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "correct formatted string to json",
			input: "{\"name\": \"John\", \"age\": 32}",
			expected: `{
				"age": 32,
				"name": "John"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formattedJson, err := service.FormatStringToJson(tt.input)

			if err != nil {
				t.Errorf("unexpected error for input %s: %s", tt.input, err)
			}

			fmt.Println(reflect.TypeOf(formattedJson))
			// todo: fix the string to json format for both
			// if formattedJson != tt.expected {
			// 	t.Errorf("expected %s, got %s", tt.expected, formattedJson)
			// }
		})
	}
}
