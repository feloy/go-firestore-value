package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	Fields interface{} `json:"fields"`
}

// GetStringValue extracts a string value from a Firestore value
func (v FirestoreValue) GetStringValue(name string) (string, error) {
	fields, ok := v.Fields.(map[string]interface{})
	mapped, ok := fields[name].(map[string]interface{})
	if !ok {
		return "", errors.New(fmt.Sprintf("Error extracting value %s from %+v", name, fields))
	}
	value, ok := mapped["stringValue"].(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("Error extracting value %s from %+v", name, fields))
	}
	return value, nil
}

// GetIntegerValue extracts an integer value from a Firestore value
func (v FirestoreValue) GetIntegerValue(name string) (int, error) {
	fields, ok := v.Fields.(map[string]interface{})
	mapped, ok := fields[name].(map[string]interface{})
	if !ok {
		return 0, errors.New(fmt.Sprintf("Error extracting value %s from %+v", name, fields))
	}
	strValue, ok := mapped["integerValue"].(string)
	if !ok {
		return 0, errors.New(fmt.Sprintf("Error extracting value %s from %+v", name, fields))
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// GetTimestampValue extracts a timestamp value from a Firestore value
func (v FirestoreValue) GetTimestampValue(name string) (time.Time, error) {
	fields, ok := v.Fields.(map[string]interface{})
	mapped, ok := fields[name].(map[string]interface{})
	if !ok {
		return time.Time{}, errors.New(fmt.Sprintf("Error extracting value %s from %+v", name, fields))
	}
	strValue, ok := mapped["timestampValue"].(string)
	if !ok {
		return time.Time{}, errors.New(fmt.Sprintf("Error extracting value %s from %+v", name, fields))
	}
	tsValue, err := time.Parse(time.RFC3339, strValue)
	return tsValue, err
}

// GetIntArrayValue returns an array of integer from a FirestoreValue
// Returns an empty array if 'name' is not a defined field
func (v FirestoreValue) GetIntArrayValue(name string) ([]int, error) {
	fields, ok := v.Fields.(map[string]interface{})
	mapped, ok := fields[name].(map[string]interface{})
	if !ok {
		return []int{}, nil
	}
	arrValue, ok := mapped["arrayValue"].(map[string]interface{})
	if !ok {
		return []int{}, errors.New(fmt.Sprintf("Error 1 extracting 'arrayValue' from %+v", mapped))
	}
	values, ok := arrValue["values"].([]interface{})
	if !ok {
		return []int{}, nil
	}
	res := make([]int, 0, len(values))
	for _, intfvalue := range values {
		value := intfvalue.(map[string]interface{})
		str, ok := value["integerValue"].(string)
		if !ok {
			return []int{}, errors.New(fmt.Sprintf("Error 3 extracting integerValue from %+v", value))
		}
		v, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, errors.New(fmt.Sprintf("Error 4 extracting integer value from %+v", str))
		}
		res = append(res, v)
	}
	return res, nil
}
