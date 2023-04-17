package contxt

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-clean-arch/pkg/ginwrapper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey int

const (
	dsbCtxKey ctxKey = iota
)

// DSBContext implements IDSBContext, and IContextLogger interface
type DSBContext struct {
	ginCtx    *gin.Context // for getter/setter implementation
	prefix    string
	reqInfo   RequestInfo
	parentCtx context.Context
}

// RequestInfo for Context
type RequestInfo struct {
	RequestID string
	ClientIP  string
	Host      string
	Method    string
	Path      string
	Referer   string
	UserAgent string
}

// NewDSBContext is used to get DSBContext out from the regular context.Context
func NewDSBContext(c context.Context) IDSBContext {
	ginWrapperCtx, ok := c.(*ginwrapper.Context)
	if ok {
		return NewDSBContext(ginWrapperCtx.Context)
	}

	ginCtx, ok := c.(*gin.Context)
	if ok {
		c = ginCtx.Request.Context()
	}
	value := c.Value(dsbCtxKey)
	requestId := getRequestID(c)
	if value == nil {
		return &DSBContext{reqInfo: RequestInfo{RequestID: requestId}}
	}

	ctx, ok := value.(*DSBContext)
	if !ok {
		return &DSBContext{reqInfo: RequestInfo{RequestID: requestId}}
	}

	ctx.parentCtx = c
	return ctx
}

func getRequestID(c context.Context) string {
	value := c.Value("RequestID")
	requestId, ok := value.(string)
	if !ok {
		// nolint
		return uuid.New().String()
	}
	return requestId
}

// NewContext is used to get DSBContext out from the regular context.Context
func NewContext(c context.Context, requestId string) IDSBContext {
	ginWrapperCtx, ok := c.(*ginwrapper.Context)
	if ok {
		return NewDSBContext(ginWrapperCtx.Context)
	}

	ginCtx, ok := c.(*gin.Context)
	if ok {
		c = ginCtx.Request.Context()
	}
	value := c.Value(dsbCtxKey)
	if value == nil {
		return &DSBContext{reqInfo: RequestInfo{RequestID: requestId}}
	}

	ctx, ok := value.(*DSBContext)
	if !ok {
		return &DSBContext{reqInfo: RequestInfo{RequestID: requestId}}
	}

	ctx.parentCtx = c
	return ctx
}

// GetLogger ...
func (c *DSBContext) GetLogger() IContextLogger {
	return c.GetLoggerWithPrefix("")
}

// GetLoggerWithPrefix ...
func (c *DSBContext) GetLoggerWithPrefix(prefix string) IContextLogger {
	return &ContextLogger{prefix: prefix, reqInfo: c.reqInfo}
}

// Set ...
func (c *DSBContext) Set(key string, value interface{}) {
	c.ginCtx.Set(key, value)
}

// Get ...
func (c *DSBContext) Get(key string) (interface{}, bool) {
	if c.ginCtx == nil {
		return nil, false
	}
	return c.ginCtx.Get(key)
}

// Request return the current request
func (c *DSBContext) Request() *http.Request {
	return c.ginCtx.Request
}

// SetUserID set user_id
func (c *DSBContext) SetUserID(userID string) {
	c.Set("user_id", userID)
}

// GetUserID return user id
func (c *DSBContext) GetUserID() (string, error) {
	value, found := c.Get("user_id")
	if !found {
		return "", errors.New("cannot get user_id from context")
	}

	userID, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("could not convert userID=%v to string", value)
	}

	return userID, nil
}

// SetDeviceID ...
func (c *DSBContext) SetDeviceID(value string) {
	c.Set("device_id", value)
}

// GetDeviceID ...
func (c *DSBContext) GetDeviceID() (string, error) {
	value, found := c.Get("device_id")
	if !found {
		return "", errors.New("cannot get device_id from context")
	}

	deviceID, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("could not convert deviceID=%v to string", value)
	}

	return deviceID, nil
}

// InternalServerError is shortcut of c.JSONError(http.StatusInternalServerError, code, message)
func (c *DSBContext) InternalServerError(err error, message string) {
	c.JSONError(http.StatusInternalServerError, "internal_server_error", message, nil)
}

// JSONResponse ...
func (c *DSBContext) JSONResponse(code int, data interface{}) {
	c.ginCtx.JSON(code, gin.H{
		"data": data,
	})
}

// InvalidDeviceToken ...
func (c *DSBContext) InvalidDeviceToken(message string, detail interface{}) {
	c.JSONError(http.StatusBadRequest, "invalid_device_token", message, detail)
	c.ginCtx.Abort()
}

// BadRequest ...
func (c *DSBContext) BadRequest(message string, detail interface{}) {
	c.JSONError(http.StatusBadRequest, "bad_request", message, detail)
	c.ginCtx.Abort()
}

// JSONData ...
func (c *DSBContext) JSONData(code int, data interface{}) {
	c.ginCtx.JSON(code, gin.H{
		"data": data,
	})
}

// JSONError response json format error
// this will Abort other handlers
func (c *DSBContext) JSONError(statusCode int, code string, message string, detail interface{}) {
	errorItem := gin.H{
		"code":    code,
		"message": message,
	}
	if detail != nil {
		errorItem["detail"] = detail
	}

	c.ginCtx.JSON(statusCode, gin.H{
		"errors": []gin.H{errorItem},
	})
	c.ginCtx.Abort()
}

// GetString ...
func (c *DSBContext) GetString(key string) (string, bool) {
	v, ok := c.Get(key)
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	if !ok {
		return "", false
	}
	return s, true
}

// GetInt64 ...
func (c *DSBContext) GetInt64(key string) (int64, bool) {
	v, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	i, ok := v.(int64)
	if !ok {
		return 0, false
	}
	return i, true
}

// GetRequestID ...
func (c *DSBContext) GetRequestID() string {
	return c.reqInfo.RequestID
}

// GetRequestInfo ...
func (c *DSBContext) GetRequestInfo() RequestInfo {
	return c.reqInfo
}

// RequestStart mark the start of the current request and
// ship request data to Stackdriver
func (c *DSBContext) RequestStart() {
	c.Set("RequestTime", time.Now())
	c.GetLogger().SetPrefix("Middleware").Infof("Request started")
}

// RequestFinished output the marker for current request and
// ship response data to Stackdriver
func (c *DSBContext) RequestFinished() {
	statusCode := c.ginCtx.Writer.Status()
	s, _ := c.Get("RequestTime")
	reqTime, _ := s.(time.Time)
	latency := int(math.Ceil(float64(time.Since(reqTime).Nanoseconds()) / 1e6))

	c.prefix = "Middleware"
	responseLength := fmt.Sprintf("%dB", c.ginCtx.Writer.Size())
	fmtLatency := fmt.Sprintf("%dms", latency)

	fields := c.basicFields()
	exfields := []zapcore.Field{
		zap.Int("statusCode", statusCode),
		zap.String("latency", fmtLatency),
		zap.String("responseLength", responseLength),
	}
	fields = append(fields, exfields...)

	// assign log level
	if statusCode > 499 {
		zap.L().Error("Request finished", fields...)
	} else if statusCode > 399 {
		zap.L().Warn("Request finished", fields...)
	} else {
		zap.L().Info("Request finished", fields...)
	}
}

// support func to build up basic fields
func (c *DSBContext) basicFields() []zapcore.Field {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	// frame of getCurrentFrame function
	frame, _ := frames.Next() // nolint
	// frame of caller of this function
	frame, _ = frames.Next() // nolint
	logLine := fmt.Sprintf("%s:%d", frame.File, frame.Line)

	fields := []zapcore.Field{
		zap.String("prefix", c.prefix),
		zap.String("request_id", c.reqInfo.RequestID),
		zap.String("client_ip", c.reqInfo.ClientIP),
		zap.String("host", c.reqInfo.Host),
		zap.String("method", c.reqInfo.Method),
		zap.String("path", c.reqInfo.Path),
		zap.String("referer", c.reqInfo.Referer),
		zap.String("user_agent", c.reqInfo.UserAgent),
		zap.String("log_line", logLine),
	}

	return fields
}
