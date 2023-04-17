package ginwrapper

import (
	"context"
	"errors"
	"fmt"
	"go-clean-arch/pkg/constant"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	xErrors "go-clean-arch/pkg/errors"
)

type ctxKey int

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

const (
	ginWrapperKey ctxKey = iota
)

// Context wrapper of gin.Context
type Context struct {
	*gin.Context
}

// ToContext ...
func (c *Context) ToContext() context.Context {
	return context.WithValue(context.Background(), ginWrapperKey, c)
}

// SetClientID ...
func (c *Context) SetClientID(clientID string) {
	c.Set("client_id", clientID)
}

// SetSellerID ...
func (c *Context) SetSellerID(sellerID string) {
	c.Set("seller_id", sellerID)
}

// SetAdminUserID set admin user
func (c *Context) SetAdminUserID(userID string) {
	c.Set("admin_user_id", userID)
}

// SetUserID set user_id
func (c *Context) SetUserID(userID string) {
	c.Set("user_id", userID)
}

// SetToken set token
func (c *Context) SetToken(token string) {
	c.Set("token", token)
}

// SetAdminUserName ...
func (c *Context) SetAdminUserName(name string) {
	c.Set("admin_user_name", name)
}

//SetIsAdmin ....
func (c *Context) SetIsAdmin(isAdmin bool) {
	c.Set("is_admin", isAdmin)
}

// GetSellerID ...
func (c *Context) GetSellerID() (string, error) {
	value, found := c.Get("seller_id")
	if !found {
		return "", errors.New("cannot get sellerid from context")
	}

	sellerID, ok := value.(string)
	if !ok {
		// nolint
		return "", fmt.Errorf("could not convert seller ID =%v to string", value)
	}

	return sellerID, nil
}

// GetUserID set user_id
func (c *Context) GetUserID() (string, error) {
	value, found := c.Get("user_id")
	if !found {
		return "", errors.New("cannot get user_id from context")
	}

	userID, ok := value.(string)
	if !ok {
		// nolint
		return "", fmt.Errorf("could not convert userID=%v to string", value)
	}

	return userID, nil
}

// GetIsAdmin ...
func (c *Context) GetIsAdmin() (bool, error) {
	value, found := c.Get("is_admin")
	if !found {
		return false, errors.New("cannot get is_admin from context")
	}

	isAdmin, ok := value.(bool)
	if !ok {
		// nolint
		return false, fmt.Errorf("could not convert value=%v to boolean", value)
	}

	return isAdmin, nil
}

// GetAdminUserID set user_id
func (c *Context) GetAdminUserID() (string, error) {
	value, found := c.Get("admin_user_id")
	if !found {
		return "", errors.New("cannot get admin_user_id from context")
	}

	userID, ok := value.(string)
	if !ok {
		// nolint
		return "", fmt.Errorf("could not convert admin userID=%v to string", value)
	}

	return userID, nil
}

// GetAdminUserName ...
func (c *Context) GetAdminUserName() (string, error) {
	value, found := c.Get("admin_user_name")
	if !found {
		return "", errors.New("cannot get admin_user_name from context")
	}

	userName, ok := value.(string)
	if !ok {
		// nolint
		return "", fmt.Errorf("could not convert admin userName=%v to string", value)
	}

	return userName, nil
}

// SetAppID set app_id
func (c *Context) SetAppID(appID string) {
	c.Set("app_id", appID)
}

// GetIPAddress returns user's IP address
func (c *Context) GetIPAddress() string {
	return c.ClientIP()
}

// GetIPAddressWithZeroLastOctet returns user's IP address with zero last octet
func (c *Context) GetIPAddressWithZeroLastOctet() string {
	ipAddress := c.ClientIP()
	index := strings.LastIndex(ipAddress, ".")
	if index == -1 {
		return ipAddress
	}

	return fmt.Sprintf("%s.%s", ipAddress[0:index], "0")
}

// GetAuthToken ...
func (c *Context) GetAuthToken() (string, error) {
	tk, ok := c.Get("token")
	if !ok {
		return "", fmt.Errorf("token not found in context")
	}

	token, ok := tk.(string)
	if !ok {
		return "", fmt.Errorf("token must be string")
	}
	return token, nil
}

// GetAuthSource ...
func (c *Context) GetAuthSource() (string, error) {
	tk, ok := c.Get("source")
	if !ok {
		return "", fmt.Errorf("token not found in context")
	}

	token, ok := tk.(string)
	if !ok {
		return "", fmt.Errorf("token must be string")
	}
	return token, nil
}

// JSONData ...
func (c *Context) JSONData(statusCode int, data interface{}) {
	c.JSONResponse(statusCode, data, "", "", "")
}

// JSONData ...
func (c *Context) PureJSONData(statusCode int, data interface{}) {
	c.PureJSONResponse(statusCode, data, "", "", "")
}

// JSONError ...
func (c *Context) JSONError(statusCode int, errorCode string, errorMessage string) {
	c.JSONResponse(statusCode, nil, errorCode, errorMessage, "")
}

// JSONResponse ...
func (c *Context) PureJSONResponse(statusCode int, data interface{}, errorCode string, errorMessage string, errorField string) {
	message := constant.MessageSuccess
	returnCode := http.StatusOK
	if statusCode == http.StatusInternalServerError {
		returnCode = http.StatusInternalServerError
	} else if statusCode == http.StatusUnauthorized {
		returnCode = http.StatusUnauthorized
	} else if statusCode == http.StatusTooManyRequests {
		returnCode = http.StatusTooManyRequests
	} else if statusCode == http.StatusBadRequest {
		returnCode = http.StatusBadRequest
	}
	if errorMessage != "" {
		message = ""
	}
	if returnCode != http.StatusOK {
		message = ""
	}

	resp := &APIResponse{}
	resp.Code = returnCode
	resp.Message = message

	if data != nil {
		resp.Data = data
	}

	if errorCode != "" && errorMessage != "" {
		resp.Errors = []gin.H{
			{
				"message": errorMessage,
				"extensions": gin.H{
					"status_code": statusCode,
					"code":        errorCode,
					"field":       errorField,
				},
			},
		}
	}

	c.PureJSON(returnCode, resp)

	if statusCode != http.StatusOK {
		c.Abort()
	}
}

// PureErrorJSONResponse ...
func (c *Context) PureErrorJSONResponse(statusCode int, error interface{}) {
	resp := gin.H{}
	if error != nil {
		resp["code"] = statusCode
		resp["errors"] = error
	}
	c.PureJSON(statusCode, resp)
}

// JSONResponse ...
func (c *Context) JSONResponse(statusCode int, data interface{}, errorCode string, errorMessage string, errorField string) {
	message := constant.MessageSuccess
	returnCode := http.StatusOK
	if statusCode == http.StatusInternalServerError {
		returnCode = http.StatusInternalServerError
	} else if statusCode == http.StatusUnauthorized {
		returnCode = http.StatusUnauthorized
	} else if statusCode == http.StatusTooManyRequests {
		returnCode = http.StatusTooManyRequests
	} else if statusCode == http.StatusBadRequest {
		returnCode = http.StatusBadRequest
	}
	if errorMessage != "" {
		message = ""
	}
	if returnCode != http.StatusOK {
		message = ""
	}

	resp := &APIResponse{}
	resp.Code = returnCode
	resp.Message = message

	if data != nil {
		resp.Data = data
	}

	if errorCode != "" && errorMessage != "" {
		resp.Errors = []gin.H{
			{
				"message": errorMessage,
				"extensions": gin.H{
					"status_code": statusCode,
					"code":        errorCode,
					"field":       errorField,
				},
			},
		}
	}

	c.JSON(returnCode, resp)

	if statusCode != http.StatusOK {
		c.Abort()
	}
}

// BadRequest ...
func (c *Context) BadRequest(err error) {
	invalidInputError, ok := err.(*xErrors.InvalidInputError)
	if ok {
		c.JSONResponse(http.StatusBadRequest, nil, "bad_input_error", invalidInputError.Message, invalidInputError.Field)
	} else {
		c.JSONResponse(http.StatusBadRequest, nil, "bad_request", err.Error(), "")
	}
}

// InvalidToken ...
func (c *Context) InvalidToken(kind string, message string) {
	code := fmt.Sprintf("invalid_%s", kind)
	c.JSONError(http.StatusBadRequest, code, message)
}

// InternalServerError is shortcut of c.JSONError(http.StatusInternalServerError, code, message)
func (c *Context) InternalServerError(err error, message string) {
	// sentry.WithScope(func(scope *sentry.Scope) {
	// 	sentry.CaptureException(err)
	// })
	//newrelic.RecordError(c, err)
	c.JSONError(http.StatusInternalServerError, "internal_server_error", message)
}

// InvalidParamError is shortcut of c.JSONError(http.StatusBadRequest, code, message)
func (c *Context) InvalidParamError(message string, detail interface{}) {
	c.JSONError(http.StatusBadRequest, "invalid_parameter_exception", message)
}

// Unauthorized is shortcut of c.JSONError(http.StatusUnauthorized, code, message)
func (c *Context) Unauthorized(code, message string) {
	c.JSONError(http.StatusUnauthorized, code, message)
}

// TooManyRequests is shortcut of c.JSONError(http.StatusUnauthorized, code, message)
func (c *Context) TooManyRequests(code, message string) {
	c.JSONError(http.StatusTooManyRequests, code, message)
}

// AccessDenied is shortcut of c.Unauthorized("access_denied_exception", "You are not authorized to perform the action")
func (c *Context) AccessDenied() {
	c.Unauthorized("access_denied_exception", "You are not authorized to perform the action")
}

// TokenExpired returns 401 with error token_expired_exception
func (c *Context) TokenExpired() {
	c.Unauthorized("token_expired_exception", "The provided token is expired")
}

// InvalidCredential returns invalid credential exception
func (c *Context) InvalidCredential() {
	c.Unauthorized("invalid_credential_exception", "Username or password is incorrect")
}

// InvalidPermission returns invalid credential exception
func (c *Context) InvalidPermission() {
	c.Unauthorized("invalid_permission_request", "permission invalid")
}

// HandlerFunc handler function that use ginwrapper.Context
type HandlerFunc func(ctx *Context)

// WithContext HOC for wrapping the gin context
func WithContext(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wrappedContext := &Context{
			ctx,
		}
		handler(wrappedContext)
	}
}

// ContextWithGinWrapper create context from ginwrapper
func ContextWithGinWrapper(ctx context.Context, wrapperCtx *Context) context.Context {
	// nolint
	return context.WithValue(ctx, ginWrapperKey, wrapperCtx)
}

// GetGinWrapper return ginwrapper from context
func GetGinWrapper(ctx context.Context) (*Context, error) {
	value := ctx.Value(ginWrapperKey)
	if value == nil {
		return nil, errors.New("could not get ginwrapper.Context from context")
	}

	wrapperCtx, ok := value.(*Context)
	if !ok {
		return nil, errors.New("could not get ginwrapper.Context from context")
	}

	return wrapperCtx, nil
}
