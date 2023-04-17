package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin/binding"
)

const contentTypeField = "Content-Type"

type builder struct {
	context     context.Context
	method      string
	url         string
	contentType string
	headers     map[string]string
	bodyData    interface{}
}

func NewRequestBuilder() *builder {
	headers := make(map[string]string)
	ctx := context.Background()
	return &builder{context: ctx, headers: headers}
}

func NewRequestBuilderWithCtx(ctx context.Context) *builder {
	headers := make(map[string]string)
	return &builder{context: ctx, headers: headers}
}

func (b *builder) WithMethod(method string) *builder {
	b.method = method
	return b
}

func (b *builder) WithBody(contentType string, data interface{}) *builder {
	b.contentType = contentType
	b.bodyData = data
	b.headers[contentTypeField] = contentType
	return b
}

func (b *builder) WithHeaders(headers map[string]string) *builder {
	for k, v := range headers {
		b.headers[k] = v
	}
	return b
}
func (b *builder) WithURL(url string) *builder {
	b.url = url
	return b
}
func (b *builder) WithContext(c context.Context) *builder {
	b.context = c
	return b
}

func (b *builder) Build() (req *http.Request, err error) {
	if b.method == http.MethodGet {
		return b.buildGetRequest()
	}
	bodyByte, err := b.buildBody()
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequestWithContext(b.context, b.method, b.url, bytes.NewReader(bodyByte))
	if err != nil {
		return nil, err
	}
	for k, v := range b.headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func (b *builder) buildGetRequest() (req *http.Request, err error) {
	req, err = http.NewRequestWithContext(b.context, b.method, b.url, nil)
	if err != nil {
		return
	}
	for k, v := range b.headers {
		req.Header.Set(k, v)
	}
	return
}

func (b *builder) buildBody() ([]byte, error) {
	switch b.contentType {
	case binding.MIMEJSON:
		return json.Marshal(b.bodyData)
	case binding.MIMEPOSTForm:
		var data string
		switch b.bodyData.(type) {
		case string:
			data = b.bodyData.(string)
		case url.Values:
			data = b.bodyData.(url.Values).Encode()
		}
		return []byte(data), nil
	case binding.MIMEMultipartPOSTForm:
		switch params := b.bodyData.(type) {
		case map[string]string:
			path := ""
			for key, val := range params {
				if key == "file" {
					path = val
				}
			}
			file, err := os.Open(path)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", filepath.Base(path))
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(part, file)
			for key, val := range params {
				if key != "file" {
					_ = writer.WriteField(key, val)
				}
			}
			err = writer.Close()
			if err != nil {
				return nil, err
			}
			b.headers["Content-Type"] = writer.FormDataContentType()
			return body.Bytes(), nil
		default:
			return nil, nil
		}
	default:
		return json.Marshal(b.bodyData)
	}
}
