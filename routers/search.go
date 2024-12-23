package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/handlers"
)

func GsheetRoutes(router *gin.Engine) {
	router.GET("/gsheet", handlers.GetSheetTitle)
	router.GET("/search", handlers.SearchValueInSheet)
}
