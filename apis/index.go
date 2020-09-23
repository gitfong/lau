package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index index page
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/index.html", gin.H{
		"title":   "lau",
		"content": "welcome",
	})
}
