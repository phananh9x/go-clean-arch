package contxt

import (
	"net/http"
)

// IDSBContext is an interface for DSBContext
type IDSBContext interface {
	GetLogger() IContextLogger
	GetLoggerWithPrefix(prefix string) IContextLogger
	GetUserID() (string, error)
	Get(key string) (val interface{}, ok bool)
	GetString(key string) (val string, ok bool)
	GetInt64(key string) (val int64, ok bool)
	GetRequestID() string
	GetRequestInfo() RequestInfo
	GetDeviceID() (string, error)

	BadRequest(message string, detail interface{})

	Request() *http.Request
	RequestStart()
	RequestFinished()

	Set(key string, val interface{})
	SetDeviceID(value string)
	SetUserID(userID string)

	InternalServerError(err error, message string)
	InvalidDeviceToken(message string, detail interface{})

	JSONResponse(code int, data interface{})
}

// IContextLogger is a minimal interface for writing logs
type IContextLogger interface {
	SetPrefix(prefix string) IContextLogger
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

// JSONResult ...
type JSONResult struct {
	Data   interface{} `json:"data"`
	Errors []Errors    `json:"errors"`
}

// Extensions ...
type Extensions struct {
	Code       string `json:"code"`
	Field      string `json:"field"`
	StatusCode int    `json:"status_code"`
}

// Errors ...
type Errors struct {
	Extensions Extensions `json:"extensions"`
	Message    string     `json:"message"`
}
