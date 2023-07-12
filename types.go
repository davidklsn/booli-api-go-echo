package main

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	UserID       string `gorm:"unique"`
	Residences   datatypes.JSON
	ActivityData datatypes.JSON
	CustomInfo   datatypes.JSON
}

type Request struct {
	ID        string                 `json:"id"`
	Residence map[string]interface{} `json:"residence"`
	Activity  map[string]interface{} `json:"activity"`
	Info      map[string]interface{} `json:"info"`
}
