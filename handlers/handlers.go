package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davidklsn/booli-api-go/types"

	"github.com/davidklsn/booli-api-go/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	HOSTNAME := os.Getenv("MYSQL_HOSTNAME")
	PASSWORD := os.Getenv("MYSQL_PASSWORD")
	USER := os.Getenv("MYSQL_USER")

	var err error
	dsn := USER + ":" + PASSWORD + "@tcp(" + HOSTNAME + ":" + MYSQL_PORT + ")/booli?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: %s", err.Error())
		panic("failed to connect database")
	}

	DB.AutoMigrate(&types.UserData{})
}

// Get all users
func GetUsers() ([]types.UserData, error) {
	var usersData []types.UserData
	result := DB.Find(&usersData)

	if result.Error != nil {
		return nil, result.Error
	}

	return usersData, nil
}

// Retrieve a user from the database given its ID
func GetUser(id string) (types.UserData, error) {
	var userData types.UserData
	result := DB.First(&userData, "user_id = ?", id)

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}

// Create new user
func CreateUser(id string, residence map[string]interface{}, activity map[string]interface{}, info map[string]interface{}) (types.UserData, error) {
	residenceArray := []map[string]interface{}{residence}

	residencesJSON, err := json.Marshal(residenceArray)
	if err != nil {
		return types.UserData{}, err
	}

	infoJSON, err := json.Marshal(info)
	if err != nil {
		return types.UserData{}, err
	}

	activityJSON, err := json.Marshal(activity)
	if err != nil {
		return types.UserData{}, err
	}

	userData := types.UserData{
		UserID:       id,
		Residences:   residencesJSON,
		ActivityData: activityJSON,
		CustomInfo:   infoJSON,
	}

	result := DB.Create(&userData)

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}

func UpdateResidences(id string, residence map[string]interface{}) (types.UserData, error) {
	var userData types.UserData
	if err := DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	var existingResidences []map[string]interface{}

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
	result := DB.Model(&userData).Updates(types.UserData{
		Residences: updatedResidencesJSON,
	})

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}

func UpdateUser(id string, residence map[string]interface{}, activity map[string]interface{}, info map[string]interface{}) (types.UserData, error) {
	var userData types.UserData
	if err := DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	var existingActivity map[string]interface{}

	// Unmarshal existing data into map
	if err := json.Unmarshal(userData.ActivityData, &existingActivity); err != nil {
		return types.UserData{}, err
	}

	// Merge the two maps, activity overwrites existingActivity
	for key, value := range activity {
		existingActivity[key] = value
	}

	// Marshal the merged map back into JSON
	updatedActivity, err := json.Marshal(existingActivity)
	if err != nil {
		return types.UserData{}, err
	}

	// Update the data in the database
	result := DB.Model(&userData).Updates(types.UserData{
		ActivityData: updatedActivity,
	})

	if result.Error != nil {
		return userData, result.Error
	}

	return userData, nil
}

// Delete exisiting user
func DeleteUser(id string) (types.UserData, error) {
	var userData types.UserData
	if err := DB.First(&userData, "user_id = ?", id).Error; err != nil {
		return types.UserData{}, err
	}

	if err := DB.Unscoped().Delete(&userData).Error; err != nil {
		return types.UserData{}, err
	}

	return userData, nil
}
