package contxt

import (
	"context"
	"fmt"
	"runtime"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GRPCLogger ...
type GRPCLogger struct {
	prefix  string
	fields  []zapcore.Field
	context context.Context
}

//NewGRPCLogger ...
func NewGRPCLogger(ctx context.Context, prefix string) IContextLogger {
	return &GRPCLogger{prefix: prefix, context: ctx}
}

//Context ...
func (l *GRPCLogger) Context() context.Context {
	return l.context
}

// Debugf ...
func (l *GRPCLogger) Debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	ctxzap.Extract(l.context).Debug(str, l.basicFields()...)
}

// Infof ...
func (l *GRPCLogger) Infof(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	ctxzap.Extract(l.context).Info(str, l.basicFields()...)
}

// Warnf ...
func (l *GRPCLogger) Warnf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	ctxzap.Extract(l.context).Warn(str, l.basicFields()...)
}

// Errorf ...
func (l *GRPCLogger) Errorf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	ctxzap.Extract(l.context).Error(str, l.basicFields()...)
}

// SetPrefix ...
func (l *GRPCLogger) SetPrefix(prefix string) IContextLogger {
	l.prefix = prefix
	return l
}

// support func to build up basic fields
func (l *GRPCLogger) basicFields() []zapcore.Field {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	// frame of getCurrentFrame function
	frame, _ := frames.Next() //nolint
	// frame of caller of this function
	frame, _ = frames.Next()
	logLine := fmt.Sprintf("%s:%d", frame.File, frame.Line)
	fields := []zapcore.Field{
		zap.String("prefix", l.prefix),
		zap.String("log_line", logLine),
	}

	return append(l.fields, fields...)
}
