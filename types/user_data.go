package types

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	UserID     string `gorm:"unique"`
	Residences datatypes.JSON
	Info       datatypes.JSON
}
