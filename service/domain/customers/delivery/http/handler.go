package http

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/pkg/contxt"
	"go-clean-arch/pkg/ginwrapper"
	"go-clean-arch/service/domain/customers/usecase"
	"net/http"
)

type customerHandler struct {
	customerUseCase usecase.ICustomerUseCase
}

func NewCustomerHandler(customerUseCase usecase.ICustomerUseCase) *customerHandler {
	return &customerHandler{
		customerUseCase: customerUseCase,
	}
}

func (h *customerHandler) GetCustomerInfo() gin.HandlerFunc {
	return ginwrapper.WithContext(func(ctx *ginwrapper.Context) {
		log := contxt.NewDSBContext(ctx).GetLoggerWithPrefix("GetCustomerInfo")
		userID, err := ctx.GetUserID()
		if err != nil {
			log.Errorf("error while GetUserID, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		resp, err := h.customerUseCase.GetCustomerInfo(ctx, userID)
		if err != nil {
			log.Errorf("error while GetCustomerInfo, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		ctx.JSONData(http.StatusOK, resp)
	})
}

func (h *customerHandler) Login() gin.HandlerFunc {
	return ginwrapper.WithContext(func(ctx *ginwrapper.Context) {
		log := contxt.NewDSBContext(ctx).GetLoggerWithPrefix("Login")

		resp, err := h.customerUseCase.GetAccessToken(ctx, ctx.Query("username"), ctx.Query("password"))
		if err != nil {
			log.Errorf("error while GetCustomerInfo, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		ctx.JSONData(http.StatusOK, resp)
	})
}
