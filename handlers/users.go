package handlers

import (
	"encoding/json"

	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/types"
)

// Get all users
func GetUsers() ([]types.UserData, error) {
	var usersData []types.UserData
	result := constants.DB.Find(&usersData)

	if result.Error != nil {
		return nil, result.Error
	}

	return usersData, nil
}

// Retrieve a user from the database given its ID
func GetUser(id string) (types.UserData, error) {
	var userData types.UserData
	result := constants.DB.First(&userData, "user_id = ?", id)

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}

// Create new user
func CreateUser(id string, residence map[string]any, activity map[string]any, info map[string]any) (types.UserData, error) {

	// RESIDENCE
	var residenceArray any
	if residence == nil {
		residenceArray = make([]map[string]any, 0)
	} else {
		residenceArray = []map[string]any{residence}
	}
	residencesJSON, err := json.Marshal(residenceArray)

	if err != nil {
		return types.UserData{}, err
	}

	// ACTIVITY 
	activityJSON, err := json.Marshal(activity)
	if err != nil {
		return types.UserData{}, err
	}

	// INFO 
	infoJSON, err := json.Marshal(info)
	if err != nil {
		return types.UserData{}, err
	}

	userData := types.UserData{
		UserID:       id,
		Residences:   residencesJSON,
		ActivityData: activityJSON,
		CustomInfo:   infoJSON,
	}

	result := constants.DB.Create(&userData)

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}

// Delete exisiting user
func DeleteUser(id string) (types.UserData, error) {
	var userData types.UserData
	if err := constants.DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	if err := constants.DB.Unscoped().Delete(&userData).Error; err != nil {
		return types.UserData{}, err
	}

	return userData, nil
}
