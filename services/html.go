package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "./src/index.html", gin.H{"title": "Main Page"})
}
