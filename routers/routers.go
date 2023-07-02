package routers

import (
	"shoppingcart.com/controllers"
	"shoppingcart.com/handlers"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	r.POST("/bill", handlers.NewHandlerWithBody(controllers.CreateBill))

	return r
}