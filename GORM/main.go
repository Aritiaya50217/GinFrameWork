package main

import (
	"example/gorm-api/config"
	"example/gorm-api/models"
	"example/gorm-api/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

// ref : http://www.lungmaker.com/go-programming/gin-mysql/
func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
