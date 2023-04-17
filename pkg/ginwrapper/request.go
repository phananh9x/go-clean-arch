package ginwrapper

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	hiddenContent         = "<HIDDEN>"
	ignoreContent         = "<IGNORE>"
	emptyContentTag       = "<EMPTY>"
	contentSizeLimitation = 10000
)

type DeviceInfo struct {
	Lat        string
	Long       string
	DeviceId   string
	AppVersion string
	ClientIP   string
	Platform   string
}

func isIgnoreRequestBody(ctx *gin.Context) bool {
	contentSize := ctx.Request.ContentLength
	// Ingore content too large
	if contentSize == -1 || contentSize >= contentSizeLimitation {
		return true
	}

	contentType := ctx.ContentType()
	// Ignore if content type is multipart form upload
	return contentType == gin.MIMEMultipartPOSTForm
}

// GetRequestBody return content of request for logging.
// Return "<HIDDEN>" incase body too large or multipart form.
func GetRequestBody(ctx *gin.Context) string {
	requestBody := hiddenContent

	if isIgnoreRequestBody(ctx) {
		return requestBody
	}

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		zap.S().With("err- ", err).Error("can't read body content from request")
		return ignoreContent
	}
	readCloser1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// We have to create a new Buffer and transfer it to request body again, because readCloser1 will be read.
	readCloser2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	ctx.Request.Body = readCloser2

	// Convert readCloser1 to String
	bytesBuffer := new(bytes.Buffer)
	_, err = bytesBuffer.ReadFrom(readCloser1)
	if err != nil {
		zap.S().With("err- ", err).Error("can't read byte array from reader")
		return ignoreContent
	}
	requestBody = bytesBuffer.String()
	if requestBody == "" {
		// Return Tag to easy filter
		return emptyContentTag
	}
	return requestBody
}

func GetDeviceInfo(ctx *gin.Context) DeviceInfo {
	var deviceInfo DeviceInfo
	headers := ctx.GetHeader("x-device-info")
	if headers != "" {
		arrHeaders := strings.Split(headers, ";")
		for _, item := range arrHeaders {
			if item != "" {
				headerValue := strings.Split(item, "=")
				if len(headerValue) == 2 {
					key := strings.ToLower(strings.TrimSpace(headerValue[0]))
					value := strings.ToLower(strings.TrimSpace(headerValue[1]))
					switch key {
					case "lat":
						deviceInfo.Lat = value
					case "long":
						deviceInfo.Long = value
					case "device-id":
						deviceInfo.DeviceId = value
					case "app-version":
						deviceInfo.AppVersion = value
					case "client-ip":
						deviceInfo.ClientIP = value
					}
				}

			}

		}
	}
	return deviceInfo
}