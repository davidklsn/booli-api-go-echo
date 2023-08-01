package helpers

import (
	"errors"
)

func UpdateActivityData(existingActivities *map[string]any, activity map[string]any) error {
	// Get key and value from activity
	key, value := getKeyValueFromJSON(&activity)

	// Check if key is empty
	if key == "" {
		return errors.New("Key is empty")
	}

	// Update or add value for key
	if castedValue, ok := value.(map[string]any); ok {
		(*existingActivities)[key] = castedValue
	} else {
		return errors.New("JSON for activity is not formatted correctly")
	}

	return nil
}
