package main

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	UserID      string `gorm:"unique"`
	ResidenceID string
	Meta        datatypes.JSON
}

type Request struct {
	ID          string                 `json:"id"`
	ResidenceID string                 `json:"residenceId"`
	Meta        map[string]interface{} `json:"meta"`
}
