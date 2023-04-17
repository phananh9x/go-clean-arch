package contxt

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//NewContextByRequestInfo ...
func NewContextByRequestInfo(req *RequestInfo) context.Context {
	var dsbContext *DSBContext
	if req == nil || req.RequestID == "" {
		dsbContext = &DSBContext{reqInfo: RequestInfo{RequestID: uuid.New().String()}}
	} else {
		dsbContext = &DSBContext{reqInfo: *req}
	}

	return context.WithValue(context.Background(), dsbCtxKey, dsbContext)
}

// SetupDSBContext is a middleware to embbed this Context type into gin.Context
func SetupDSBContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), dsbCtxKey, initContext(c))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// NewTestContext return a test Context with no request info ONLY USE for testing
func NewTestContext() *DSBContext {
	c := &gin.Context{}
	var requestID, prefix string
	if reqID, ok := c.Get("RequestID"); ok {
		requestID, _ = reqID.(string)
	}
	if px, ok := c.Get("LogPrefix"); ok {
		prefix, _ = px.(string)
	}

	return &DSBContext{
		ginCtx:  c,
		prefix:  prefix,
		reqInfo: RequestInfo{RequestID: requestID},
	}
}
