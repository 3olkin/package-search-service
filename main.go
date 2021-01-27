package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
	"github.com/z4rx/search-service/Config"
	"github.com/z4rx/search-service/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	r := Routes.SetupRouter()
	r.Run()
}
