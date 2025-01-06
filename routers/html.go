package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/services"
)

func RegisterStaticRoutes(router *gin.Engine) {
	/* CSS & JS */
	router.Static("/dist", "./src/dist")
	/* HTML */
	router.LoadHTMLGlob("./src/*.html")
}

func RegisterHTMLRoutes(router *gin.Engine) {
	router.GET("/", services.Main_Page)
}
