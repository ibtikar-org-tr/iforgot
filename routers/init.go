package routers

import (
	"github.com/gin-gonic/gin"
)

func AllRouters() {
	router := gin.New()

	GsheetRoutes(router)

	router.Run(":8012")
}
