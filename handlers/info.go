package handlers

import (
	"encoding/json"

	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/types"
)

func UpdateInfo(id string, info map[string]any) (types.UserData, error) {
	var userData types.UserData

	if err := constants.DB.First(&userData, "user_id = ?", id).Select("info").Error; err != nil {
		return types.UserData{}, err
	}

	var existingInfo map[string]any

	// Unmarshal existing data into slice of maps
	if err := json.Unmarshal(userData.Info, &existingInfo); err != nil {
		return types.UserData{}, err
	}

	if existingInfo != nil {
		for key := range info {
			existingInfo[key] = info[key]
		}
	} else {
		existingInfo = info
	}
	
	// Marshal updated data back into JSON
	existingInfoJSON, err := json.Marshal(existingInfo)
	if err != nil {
		return types.UserData{}, err
	}

	// Update the data in the database
	result := constants.DB.Model(&userData).Updates(types.UserData{
		Info: existingInfoJSON,
	})

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}
