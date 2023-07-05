package routes

import (
	"SpaceTradersAgent/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/status", controllers.Status)


}