package helpers

import (
	"errors"
)

func UpdateActivityData(existingActivities map[string]map[string]any, activity map[string]any) (map[string]map[string]any, error) {
	// Get key and value from activity
	key, value := getKeyValueFromJSON(activity)

	// Check if key is empty
	if key == "" {
		return nil, errors.New("Key is empty")
	}

	// Update or add value for key
	if castedValue, ok := value.(map[string]interface{}); ok {
		existingActivities[key] = castedValue
	} else {
		return nil, errors.New("JSON for activity is not formatted correctly")
	}


	return existingActivities, nil
}
