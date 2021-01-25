package main

import (
	"fmt"
	"search-service/Config"
	"search-service/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	r := Routes.SetupRouter()
	r.Run()
}
