package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/services"
)

func GetSheetTitle(c *gin.Context) {
	sheetID := c.Query("sheetID")
	sheetName, err := services.GetSheetTitle(sheetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sheetName": sheetName})
}

func SearchValueInSheet(c *gin.Context) {
	sheetID := c.Query("sheetID")
	valueToSearch := c.Query("value")
	pageName := c.Query("page")
	lastColumn := c.Query("last")

	searchResult, err := services.SearchValueInSheet(valueToSearch, sheetID, pageName, lastColumn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"searchResult": searchResult})
}
