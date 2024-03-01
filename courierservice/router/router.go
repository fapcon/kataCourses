package router

import (
	"courierservice/module/courierfacade/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	courier *controller.CourierController
}

func NewRouter(courier *controller.CourierController) *Router {
	return &Router{courier: courier}
}

func (r *Router) CourierAPI(router *gin.RouterGroup) {
	// прописать роуты для courier API
	router.GET("/status", r.courier.GetStatus)
	router.GET("/ws", r.courier.Websocket)
}

func (r *Router) Swagger(router *gin.RouterGroup) {
	router.GET("/swagger", swaggerUI)
	router.GET("/swagger/api/status", r.courier.GetStatus)
	//router.StaticFile("/docs/*", "/app/public/swagger.json")
}
