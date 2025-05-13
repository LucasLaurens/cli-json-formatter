package service

import (
	"encoding/json"
	"fmt"
)

func FormatStringToJson(value string) (string, error) {
	mappedKeyWithValue, err := mapKeyWithValue(&value)
	if err != nil {
		panic(err)
	}

	indentMappedKeyValuePair, err := indentMappedKeyValuePair(&mappedKeyWithValue)
	if err != nil {
		panic(err)
	}

	return string(indentMappedKeyValuePair), nil
}

func indentMappedKeyValuePair(formattedJson *map[string]interface{}) ([]byte, error) {
	indentedJson, err := json.MarshalIndent(*formattedJson, "", "  ")
	if err != nil {
		return nil, fmt.Errorf(
			"you are not allowed to indent the json: %v",
			err,
		)
	}

	return indentedJson, nil
}

func mapKeyWithValue(value *string) (map[string]interface{}, error) {
	var mappedKeyWithValue map[string]interface{}
	err := json.Unmarshal([]byte(*value), &mappedKeyWithValue)
	if err != nil {
		return nil, fmt.Errorf(
			"you cannot parse this string to json format: %v",
			err,
		)
	}

	return mappedKeyWithValue, nil
}
