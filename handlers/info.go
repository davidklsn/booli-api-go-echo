package handlers

import (
	"encoding/json"

	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/types"
)

func UpdateInfo(id string, info map[string]any) (types.UserData, error) {
	var userData types.UserData

	if err := constants.DB.First(&userData, "user_id = ?", id).Select("custom_info").Error; err != nil {
		return types.UserData{}, err
	}

	var existingInfo map[string]any

	// Unmarshal existing data into slice of maps
	if err := json.Unmarshal(userData.CustomInfo, &existingInfo); err != nil {
		return types.UserData{}, err
	}

	for key := range info {
		existingInfo[key] = info[key]
	}

	// Marshal updated data back into JSON
	existingInfoJSON, err := json.Marshal(existingInfo)
	if err != nil {
		return types.UserData{}, err
	}

	// Update the data in the database
	result := constants.DB.Model(&userData).Updates(types.UserData{
		CustomInfo: existingInfoJSON,
	})

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}
