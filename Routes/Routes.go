package Routes

import (
	"github.com/gin-gonic/gin"

	"github.com/z4rx/search-service/Controllers"
)

// Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/srch/api")
	{
		grp1.GET("/distros", Controllers.GetDistros)
		grp1.GET("/cli-distros", Controllers.GetCliDistros)
		grp1.GET("/parent/:distro", Controllers.GetDistroParent)
	}
	grp2 := r.Group("/srch/api/v1")
	{
		grp2.GET("/:package/:distro", Controllers.GetPackage)
	}
	return r
}
