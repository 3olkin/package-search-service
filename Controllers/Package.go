package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/z4rx/search-service/Models"
)

func GetPackage(c *gin.Context) {
	var distro Models.Distro
	err := Models.GetDistroByName(&distro, c.Params.ByName("distro"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "We don't have yet any information about this distribution.", "data": nil})
		return
	}

	var pkg Models.Package
	err = Models.GetPackageByName(&pkg, c.Params.ByName("package"), strconv.Itoa(distro.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "We couldn't find the given package.", "data": nil})
		return
	}

	var matching Models.Matching
	err = Models.GetMatchingByParams(&matching, strconv.Itoa(pkg.ID), strconv.Itoa(distro.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "We couldn't find the matching package for your distribution.", "data": nil})
		return
	}

	var result Models.Package
	err = Models.GetPackageByID(&result, strconv.Itoa(matching.MatchingID))
	// Low chance but possible error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Provide us any info about your search query if you see this error.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": nil, "data": result.Name})
}
