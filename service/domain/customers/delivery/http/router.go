package http

import "github.com/gin-gonic/gin"

//CustomerAPIRoute ...
func (h *customerHandler) CustomerAPIRoute(router *gin.RouterGroup) {
	router.GET("/customers/info", h.GetCustomerInfo())
}

//CustomerAPIRoutePublic ...
func (h *customerHandler) CustomerAPIRoutePublic(router *gin.RouterGroup) {
	router.POST("/customers/login", h.Login())
	router.POST("/customers/signup", h.Signup())
}
