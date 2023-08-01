package handlers

import (
	"encoding/json"

	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/helpers"
	"github.com/davidklsn/booli-api-go/types"
)

func UpdateActivites(id string, activity map[string]any) (types.UserData, error) {
	var userData types.UserData
	if err := constants.DB.First(&userData, "user_id = ?", id).Select("ActivityData").Error; err != nil {
		return types.UserData{}, err
	}

	var existingActivities map[string]any

	// Unmarshal existing data into slice of maps
	if err := json.Unmarshal(userData.ActivityData, &existingActivities); err != nil {
		return types.UserData{}, err
	}

	if existingActivities != nil {
		helpers.UpdateActivityData(&existingActivities, activity)
	} else {
		existingActivities = make(map[string]any)
		existingActivities = activity
	}

	updatedActivitiesJSON, err := json.Marshal(existingActivities)
	if err != nil {
		return types.UserData{}, err
	}

	// Update the data in the database
	result := constants.DB.Model(&userData).Updates(types.UserData{
		ActivityData: updatedActivitiesJSON,
	})

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}
