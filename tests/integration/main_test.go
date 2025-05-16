package tests

import (
	"cli-json-formatter/internal/service"
	"encoding/json"
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

			var expectedObj, resultObj map[string]interface{}

			if err := json.Unmarshal([]byte(tt.expected), &expectedObj); err != nil {
				t.Fatalf("invalid expected JSON: %s", err)
			}
			if err := json.Unmarshal([]byte(formattedJson), &resultObj); err != nil {
				t.Fatalf("invalid formatted JSON from function: %s", err)
			}

			if !reflect.DeepEqual(expectedObj, resultObj) {
				t.Errorf("expected %+v, got %+v", expectedObj, resultObj)
			}
		})
	}
}
