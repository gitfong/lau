package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitfong/lau/models"
)

// Author get author's info
func Author(c *gin.Context) {
	a := models.Author{}
	err := a.LoadInfo()
	if err != nil {
		fmt.Printf("LoadInfo err %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"name":     a.Name,
		"city":     a.City,
		"province": a.Province,
	})
}
