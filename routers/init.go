package routers

import (
	"github.com/gin-gonic/gin"
)

func AllRouters() {
	router := gin.New()

	RegisterStaticRoutes(router)
	RegisterHTMLRoutes(router)

	GsheetRoutes(router)

	router.Run(":8013")
}
