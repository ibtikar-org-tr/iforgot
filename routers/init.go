package routers

import (
	"github.com/gin-gonic/gin"
)

func AllRouters() {
	router := gin.New()

	router.Run(":8080")
}
