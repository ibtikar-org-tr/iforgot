package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/services"
)

func GetSheetTitle(c *gin.Context) {
	sheetID := c.Query("sheetID")
	sheetName := services.GetSheetTitle(sheetID)
	c.JSON(http.StatusOK, gin.H{"sheetName": sheetName})
}
