package handlers

import (
	"encoding/json"
	"errors"

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

	errors := helpers.UpdateResidenceData(&existingResidences, residence)

	if errors != nil {
		return types.UserData{}, errors
	}

	// Marshal updated data back into JSON
	updatedResidencesJSON, err := json.Marshal(existingResidences)
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

func UpdateCurrentResidence(id string, residence map[string]any) (types.UserData, error) {
	var userData types.UserData
	if err := constants.DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	var existingResidences []map[string]any

	// Unmarshal existing data into slice of maps
	if err := json.Unmarshal(userData.Residences, &existingResidences); err != nil {
		return types.UserData{}, err
	}

	errors := helpers.SetCurrentResidence(&existingResidences, residence)

	if errors != nil {
		return types.UserData{}, errors
	}

	// Marshal updated data back into JSON
	updatedResidencesJSON, err := json.Marshal(existingResidences)
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

func UpdateSelectedResidence(id string, residence map[string]any) (types.UserData, error) {
	var userData types.UserData
	if err := constants.DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	var existingResidences []map[string]any

	// Unmarshal existing data into slice of maps
	if err := json.Unmarshal(userData.Residences, &existingResidences); err != nil {
		return types.UserData{}, err
	}

	errors := helpers.SetSelectedResidence(&existingResidences, residence)

	if errors != nil {
		return types.UserData{}, errors
	}

	// Marshal updated data back into JSON
	updatedResidencesJSON, err := json.Marshal(existingResidences)
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

func GetCurrentResidence(id string) (map[string]any, error) {
	var UserData types.UserData
	var residences []map[string]any
	if err := constants.DB.First(&UserData, "user_id = ?", id).Select("residences").Error; err != nil {
		return make(map[string]any), err
	}

	if err := json.Unmarshal(UserData.Residences, &residences); err != nil {
		return make(map[string]any), err
	}

	for _, r := range residences {
		if val, ok := r["currentResidence"]; ok && val == true {
			return r, nil
		}
	}

	return map[string]any{}, errors.New("No current residence found")
}
