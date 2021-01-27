package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/z4rx/search-service/Config"
	"github.com/z4rx/search-service/Models"
	"github.com/z4rx/search-service/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Distro{}, &Models.Package{}, &Models.Matching{})
	r := Routes.SetupRouter()
	r.Run()
}
