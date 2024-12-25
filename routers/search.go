package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/handlers"
)

func GsheetRoutes(router *gin.Engine) {
	router.GET("/search", handlers.SearchHandler)
	router.GET("/sheet", handlers.GetSheetTitle)
	router.GET("/value", handlers.SearchValueInSheet)
}
