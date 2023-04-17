package contxt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ContextLogger ...
type ContextLogger struct {
	prefix  string
	reqInfo RequestInfo
}

// NewContextLoggerFromReqInfo ...
func NewContextLoggerFromReqInfo(reqInfo RequestInfo) *ContextLogger {
	return &ContextLogger{reqInfo: reqInfo}
}

// Debugf log formatted info with requestID
func (l *ContextLogger) Debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	zap.L().Debug(str, l.basicFields()...)
}

// Infof log formatted info with requestID
func (l *ContextLogger) Infof(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	zap.L().Info(str, l.basicFields()...)
}

// Warnf log formatted warn with requestID
func (l *ContextLogger) Warnf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	zap.L().Warn(str, l.basicFields()...)
}

// Errorf log formatted error with requestID
func (l *ContextLogger) Errorf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	zap.L().Error(str, l.basicFields()...)
}

// SetPrefix ...
func (l *ContextLogger) SetPrefix(prefix string) IContextLogger {
	l.prefix = prefix
	return l
}

// support func to build up basic fields
func (l *ContextLogger) basicFields() []zapcore.Field {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	// frame of getCurrentFrame function
	frame, _ := frames.Next() // nolint
	// frame of caller of this function
	frame, _ = frames.Next() // nolint
	logLine := fmt.Sprintf("%s:%d", frame.File, frame.Line)

	fields := []zapcore.Field{
		zap.String("prefix", l.prefix),
		zap.String("request_id", l.reqInfo.RequestID),
		zap.String("client_ip", l.reqInfo.ClientIP),
		zap.String("host", l.reqInfo.Host),
		zap.String("method", l.reqInfo.Method),
		zap.String("path", l.reqInfo.Path),
		zap.String("referer", l.reqInfo.Referer),
		zap.String("user_agent", l.reqInfo.UserAgent),
		zap.String("log_line", logLine),
	}

	return fields
}

// ex: field: password, regexPattern: all -> password: abc => password: ***
// ex: field: phone, regexPattern: (?P<TEST>[0-9]*)(?P<TEST>[0-9]{4}) -> phone: 0335266789 => phone: ******6789
func MaskJson(data string, field, regexPattern string) string {
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(data), &jsonMap); err != nil {
		return data
	}
	byteData, err := json.Marshal(sentitiveLogging(jsonMap, field, regexPattern))
	if err == nil {
		return string(byteData)
	}
	return data
}

func sentitiveLogging(data map[string]interface{}, fieldSentitive, fieldMask string) map[string]interface{} {
	if fieldMask == "" || fieldSentitive == "" {
		return data
	}
	for key, value := range data {
		if value == nil || value == "" {
			continue
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			temporaryMaps := value.([]interface{})
			for _, mapItem := range temporaryMaps {
				if reflect.TypeOf(mapItem).Kind() != reflect.Map {
					continue
				}
				sentitiveLogging(mapItem.(map[string]interface{}), fieldSentitive, fieldMask)
			}
		case reflect.Map:
			sentitiveLogging(value.(map[string]interface{}), fieldSentitive, fieldMask)
		default:
			data[key] = maskLogging(value, key, fieldSentitive, fieldMask)
		}
	}
	return data
}

func maskLogging(value interface{}, key, fieldSentitive, condition string) string {
	defer func() {
		_ = recover()
	}()
	if !strings.EqualFold(key, fieldSentitive) {
		return value.(string)
	}
	conditionStr := strings.ToLower(condition)
	if conditionStr == "" {
		return value.(string)
	}
	strValue := reflect.ValueOf(value).String()
	var maskValue string
	switch conditionStr {
	case "all":
		maskValue = regexp.MustCompile(".").ReplaceAllLiteralString(strValue, "*")
	default:
		re := regexp.MustCompile(condition)
		reValues := re.FindStringSubmatch(strValue)
		reNames := re.SubexpNames()
		for i := 1; i < len(reNames); i++ {
			if strings.ToLower(reNames[i]) == "sentitive" {
				maskValue += regexp.MustCompile(".").ReplaceAllLiteralString(reValues[i], "*")
			} else {
				maskValue += reValues[i]
			}
		}
	}
	return maskValue
}
