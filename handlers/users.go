package handlers

import (
	"encoding/json"
	"strings"

	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/types"
	"gorm.io/gorm"
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

func GetUsersByIds(ids ...string) ([]types.UserData, error) {
	var users []types.UserData
	var result *gorm.DB

	if len(ids) == 0 || ids[0] == "" {
		result = constants.DB.Find(&users)
	} else {
		idString := ids[0]
		idSlice := strings.Split(idString, ",")
		result = constants.DB.Where("user_id IN (?)", idSlice).Find(&users)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
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
func CreateUser(id string, residence map[string]any, info map[string]any) (types.UserData, error) {

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

	// INFO
	infoJSON, err := json.Marshal(info)
	if err != nil {
		return types.UserData{}, err
	}

	userData := types.UserData{
		UserID:     id,
		Residences: residencesJSON,
		Info:       infoJSON,
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
