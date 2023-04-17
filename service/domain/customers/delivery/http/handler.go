package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-clean-arch/pkg/contxt"
	"go-clean-arch/pkg/ginwrapper"
	"go-clean-arch/service/domain/customers/usecase"
	"go-clean-arch/service/models/dto"
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

// GetCustomerInfo ...
// @Summary GetCustomerInfo customer
// @Description GetCustomerInfo customer
// @Security JWT
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {object} models.JSONResult{data=dto.CustomerInfo}
// @Failure 400 {object} models.Errors
// @Failure 401 {object} models.Errors
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /v1/customers/info [get]
func (h *customerHandler) GetCustomerInfo() gin.HandlerFunc {
	return ginwrapper.WithContext(func(ctx *ginwrapper.Context) {
		log := contxt.NewDSBContext(ctx).GetLoggerWithPrefix("GetCustomerInfo")
		userID, err := ctx.GetUserID()
		if err != nil {
			log.Errorf("error while GetUserID, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		if userID == "" {
			log.Errorf("error while GetUserID, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, errors.New("user not found"))
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

// Login ...
// @Summary Login customer
// @Description Login customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param Body body dto.CustomerLoginRequest true "The body of login"
// @Success 200 {object} models.JSONResult{data=dto.CustomerLoginResponse}
// @Failure 400 {object} models.Errors
// @Failure 401 {object} models.Errors
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /v1/customers/login [post]
func (h *customerHandler) Login() gin.HandlerFunc {
	return ginwrapper.WithContext(func(ctx *ginwrapper.Context) {
		log := contxt.NewDSBContext(ctx).GetLoggerWithPrefix("Login")
		var req dto.CustomerLoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.Errorf("error while binding json, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		resp, err := h.customerUseCase.Login(ctx, req)
		if err != nil {
			log.Errorf("error while GetCustomerInfo, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		ctx.JSONData(http.StatusOK, resp)
	})
}

// Signup ...
// @Summary Register customer
// @Description register new customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param Body body dto.CustomerSignUpRequest true "The body of add new customer"
// @Success 200 {object} models.JSONResult{data=dto.CustomerSignUpResponse}
// @Failure 400 {object} models.Errors
// @Failure 401 {object} models.Errors
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /v1/customers/signup [post]
func (h *customerHandler) Signup() gin.HandlerFunc {
	return ginwrapper.WithContext(func(ctx *ginwrapper.Context) {
		log := contxt.NewDSBContext(ctx).GetLoggerWithPrefix("Login")
		var req dto.CustomerSignUpRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.Errorf("error while binding json, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		resp, err := h.customerUseCase.SignUp(ctx, req)
		if err != nil {
			log.Errorf("error while GetCustomerInfo, %v", err)
			ctx.PureErrorJSONResponse(http.StatusBadRequest, err)
			return
		}
		ctx.JSONData(http.StatusOK, resp)
	})
}
