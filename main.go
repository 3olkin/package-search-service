package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/z4rx/search-service/Config"
	"github.com/z4rx/search-service/Routes"
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
