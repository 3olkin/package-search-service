package Models

import (
	"github.com/z4rx/search-service/Config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetPackageByID(pkg *Package, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(pkg).Error; err != nil {
		return err
	}
	return nil
}

// ? поиск совпадающих имен
// функция для внутреннего поиска (pkgfind build-essential --opensuse)
func GetPackageByName(pkg *Package, name string, id string) (err error) {
	if err = Config.DB.Where("name = ? and distro_id <> ?", name, id).First(pkg).Error; err != nil {
		return err
	}
	return nil
}
