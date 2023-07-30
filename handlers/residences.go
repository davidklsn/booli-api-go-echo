package handlers

import (
	"encoding/json"

	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/helpers"
	"github.com/davidklsn/booli-api-go/types"
)

func UpdateResidences(id string, residence map[string]any) (types.UserData, error) {
	var userData types.UserData
	if err := constants.DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	var existingResidences []map[string]any

	// Unmarshal existing data into slice of maps
	if err := json.Unmarshal(userData.Residences, &existingResidences); err != nil {
		return types.UserData{}, err
	}

	updatedResidences := helpers.UpdateResidenceData(existingResidences, residence)

	// Marshal updated data back into JSON
	updatedResidencesJSON, err := json.Marshal(updatedResidences)
	if err != nil {
		return types.UserData{}, err
	}

	// Update the data in the database
	result := constants.DB.Model(&userData).Updates(types.UserData{
		Residences: updatedResidencesJSON,
	})

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}
