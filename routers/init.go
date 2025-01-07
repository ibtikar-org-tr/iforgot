package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/lib"
)

func AllRouters() {
	router := gin.New()

	// middleware
	router.Use(gin.Recovery(), lib.Logger())

	// register routes
	RegisterStaticRoutes(router)
	RegisterHTMLRoutes(router)

	GsheetRoutes(router)

	router.Run(":8013")
}
