package Models

import (
	"github.com/z4rx/search-service/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetMatchingByParams(matching *Matching, package_id string, distro_id string) (err error) {
	if err = Config.DB.Where("package_id = ? and distro_id = ?", package_id, distro_id).First(matching).Error; err != nil {
		return err
	}
	return nil
}
