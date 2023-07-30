package constants

import (
	"fmt"
	"os"

	"github.com/davidklsn/booli-api-go/types"

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
