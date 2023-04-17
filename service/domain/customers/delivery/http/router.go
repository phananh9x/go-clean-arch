package http

import "github.com/gin-gonic/gin"

//CustomerAPIRoute ...
func (h *customerHandler) CustomerAPIRoute(router *gin.RouterGroup) {
	router.GET("/customers/info", h.GetCustomerInfo())
}

//CustomerAPIRoutePublic ...
func (h *customerHandler) CustomerAPIRoutePublic(router *gin.RouterGroup) {
	router.GET("/customers/token", h.Login())
}
