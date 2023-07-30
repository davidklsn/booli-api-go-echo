package types

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
