package contxt

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	var (
		config zap.Config
		level  zapcore.LevelEncoder
	)

	switch os.Getenv("ENVIRONMENT") {
	case "local":
		level = zapcore.CapitalColorLevelEncoder
		config = zap.NewDevelopmentConfig()
	case "prod":
		level = zapcore.CapitalLevelEncoder
		config = zap.NewProductionConfig()
	default:
		level = zapcore.CapitalLevelEncoder
		config = zap.NewDevelopmentConfig()
		config.Encoding = "json"
	}

	config.Development = false
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "time",
		CallerKey:      "logger",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    level,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	defer logger.Sync() //nolint
	zap.ReplaceGlobals(logger)
}

// initContext creates a basic context from gin.Context
func initContext(c *gin.Context) *DSBContext {
	var requestID, prefix string
	if reqID, ok := c.Get("RequestID"); ok {
		requestID, _ = reqID.(string)
	}
	if px, ok := c.Get("LogPrefix"); ok {
		prefix, _ = px.(string)
	}

	ctx := &DSBContext{
		ginCtx:  c,
		prefix:  prefix,
		reqInfo: RequestInfo{RequestID: requestID},
	}
	if c.Request != nil {
		ctx.reqInfo = RequestInfo{
			RequestID: ctx.reqInfo.RequestID,
			ClientIP:  c.ClientIP(),
			Host:      c.Request.Host,
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Referer:   c.Request.Referer(),
			UserAgent: c.Request.UserAgent(),
		}
		ctx.parentCtx = c.Request.Context()
	}
	return ctx
}
