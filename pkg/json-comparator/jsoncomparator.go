package jsoncomparator

import (
	"encoding/json"
	"fmt"
)

type JSONComparator struct {
}

func New() JSONComparator {
	return JSONComparator{}
}

// CompareJSONKeys compares between two jsons to see whether there is a difference in missing keys.
// It compares both JSONs keys (not including the values).
// It returns string message, and boolean of whether there is a difference or not.
func (j *JSONComparator) CompareJSONKeys(json1, json2 string) (string, bool) {
	var message string
	var isDifferent bool

	// Unmarshal JSON data into maps
	var data1, data2 map[string]interface{}
	if err := json.Unmarshal([]byte(json1), &data1); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(json2), &data2); err != nil {
		panic(err)
	}

	// Compare keys in data1 against data2
	message += "Keys in data1 missing in data2:"
	for key := range data1 {
		if _, ok := data2[key]; !ok {
			message += fmt.Sprintf("%s, ", key)
			isDifferent = true
		}
	}

	// Compare keys in data2 against data1
	message += "\nKeys in data2 missing in data1:"
	for key := range data2 {
		if _, ok := data1[key]; !ok {
			message += fmt.Sprintf("%s, ", key)
			isDifferent = true
		}
	}

	return message, isDifferent
}
