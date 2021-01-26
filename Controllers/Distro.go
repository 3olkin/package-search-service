package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/z4rx/search-service/Models"
)

type cliResponse struct {
	Values []string `json:"values"`
}

func GetDistros(c *gin.Context) {
	var distros []Models.Distro
	err := Models.GetAllDistros(&distros)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Something went wrong. We couldn't find information about distributions.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": nil, "data": distros})
}

func GetCliDistros(c *gin.Context) {
	var distros []Models.Distro
	err := Models.GetAllDistros(&distros)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Something went wrong. We couldn't find information about distributions.", "data": nil})
		return
	}

	var responseData []string
	for _, distro := range distros {
		responseData = append(responseData, distro.Name)
	}
	values := cliResponse{Values: responseData}
	c.JSON(http.StatusOK, gin.H{"error": nil, "data": values})
}

func GetDistroParent(c *gin.Context) {
	var distro Models.Distro
	err := Models.GetDistroByName(&distro, c.Params.ByName("distro"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "We don't have yet any information about this distribution.", "data": nil})
		return
	}

	// Possible better solution exists
	if distro.ParentID == 0 {
		c.JSON(http.StatusOK, gin.H{"error": nil, "data": nil})
		return
	}

	var result Models.Distro
	err = Models.GetDistroByID(&result, strconv.Itoa(distro.ParentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Something went wrong. We couldn't find information about this distribution.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": nil, "data": result.Name})
}
