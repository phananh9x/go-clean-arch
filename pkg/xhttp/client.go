package xhttp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go-clean-arch/pkg/contxt"

	"github.com/google/go-querystring/query"
)

const (
	defaultTimeout       = 30 * time.Second
	defaultLogBodyLength = 3000
	defaultNamespace     = "go-clean-arch"
	defaultSubsystem     = "go-clean-arch"
)

type Client interface {
	PostJSON(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	PostForm(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	Get(c context.Context, url string, target interface{}, reqOptions ...RequestOption) (int, error)
	GetWithQuery(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	GetWithoutEncodedQuery(c context.Context,
		url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	Do(ctx context.Context, request *http.Request, target interface{}) (int, error)
	Put(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	Delete(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	PostMultipartForm(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
}

type client struct {
	client *http.Client
	opts   clientOptions
}

func NewClient(opts ...Option) Client {
	optsArg := getOptionsArg(opts)
	if optsArg.log == nil {
		optsArg.log = contxt.NewDSBContext(context.TODO()).GetLoggerWithPrefix("http_client")
	}
	transport := NewTransport(optsArg)
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   optsArg.timeout,
	}
	c := &client{
		client: httpClient,
		opts:   optsArg,
	}
	return c
}

func getOptionsArg(opts []Option) clientOptions {
	// Init default options arg
	optsArgs := clientOptions{
		skipLog:         false,
		splitLogBody:    false,
		splitLogBodyLen: defaultLogBodyLength,
		timeout:         defaultTimeout,
	}

	for _, opt := range opts {
		opt.apply(&optsArgs)
	}
	return optsArgs
}

func (c *client) PostJSON(ctx context.Context,
	url string, data, target interface{}, reqOpts ...RequestOption) (int, error) {
	var (
		statusCode int
		err        error
		req        *http.Request
	)
	header := c.getRequestHeader(reqOpts...)
	req, err = NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodPost).
		WithURL(url).
		WithHeaders(header).
		WithBody(MIMEJSON, data).
		Build()
	if err != nil {
		return statusCode, err
	}
	return c.Do(ctx, req, target)
}

func (c *client) PostForm(ctx context.Context,
	url string, data, target interface{}, reqOpts ...RequestOption) (int, error) {
	var (
		statusCode int
		err        error
		req        *http.Request
	)
	header := c.getRequestHeader(reqOpts...)
	req, err = NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodPost).
		WithURL(url).
		WithHeaders(header).
		WithBody(MIMEPOSTForm, data).
		Build()
	if err != nil {
		return statusCode, err
	}
	return c.Do(ctx, req, target)
}

func (c *client) PostMultipartForm(ctx context.Context,
	url string, data, target interface{}, reqOpts ...RequestOption) (int, error) {
	var (
		statusCode int
		err        error
		req        *http.Request
	)
	header := c.getRequestHeader(reqOpts...)
	req, err = NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodPost).
		WithURL(url).
		WithHeaders(header).
		WithBody(MIMEMultipartPOSTForm, data).
		Build()
	if err != nil {
		return statusCode, err
	}
	return c.Do(ctx, req, target)
}

func (c *client) Get(ctx context.Context,
	url string, target interface{}, reqOpts ...RequestOption) (int, error) {
	var (
		statusCode int
		err        error
		req        *http.Request
	)
	header := c.getRequestHeader(reqOpts...)
	req, err = NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodGet).
		WithURL(url).
		WithHeaders(header).
		Build()
	if err != nil {
		return statusCode, err
	}
	return c.Do(ctx, req, target)
}

func (c *client) GetWithQuery(ctx context.Context,
	url string, data, target interface{}, reqOpts ...RequestOption) (int, error) {
	var (
		statusCode int
		err        error
		req        *http.Request
	)
	header := c.getRequestHeader(reqOpts...)
	req, err = NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodGet).
		WithURL(url).
		WithHeaders(header).
		Build()
	if err != nil {
		return statusCode, err
	}

	if data != nil {
		v, err := query.Values(data)
		if err != nil {
			return 0, err
		}
		req.URL.RawQuery = v.Encode()
	}
	return c.Do(ctx, req, target)
}

func (c *client) Put(ctx context.Context,
	url string, data, target interface{}, reqOpts ...RequestOption) (statusCode int, err error) {
	header := c.getRequestHeader(reqOpts...)
	req, err := NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodPut).
		WithURL(url).
		WithHeaders(header).
		WithBody(MIMEJSON, data).
		Build()
	if err != nil {
		return
	}
	return c.Do(ctx, req, target)
}

func (c *client) Delete(ctx context.Context,
	url string, data, target interface{}, reqOpts ...RequestOption) (statusCode int, err error) {
	header := c.getRequestHeader(reqOpts...)
	req, err := NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodDelete).
		WithURL(url).
		WithHeaders(header).
		WithBody(MIMEJSON, data).
		Build()
	if err != nil {
		return
	}
	return c.Do(ctx, req, target)
}

func (c *client) GetWithoutEncodedQuery(ctx context.Context,
	reqURL string, data, target interface{}, reqOpts ...RequestOption) (int, error) {
	var (
		statusCode int
		err        error
		req        *http.Request
	)
	header := c.getRequestHeader(reqOpts...)
	req, err = NewRequestBuilderWithCtx(ctx).
		WithMethod(http.MethodGet).
		WithURL(reqURL).
		WithHeaders(header).
		Build()
	if err != nil {
		return statusCode, err
	}

	if data != nil {
		v, err := query.Values(data)
		if err != nil {
			return 0, err
		}
		nonEncodedValue, _ := url.QueryUnescape(v.Encode())
		req.URL.RawQuery = nonEncodedValue
	}
	return c.Do(ctx, req, target)
}

func (c *client) Do(ctx context.Context, request *http.Request, target interface{}) (int, error) {
	if requestID := request.Header.Get(RequestIDHeader); requestID == "" {
		request.Header.Set(RequestIDHeader, getContextIDFromCtx(ctx))
	}
	rsp, err := c.client.Do(request)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rsp.Body.Close()
	}()

	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}

	if len(bodyBytes) == 0 {
		return rsp.StatusCode, nil
	}

	return rsp.StatusCode, json.Unmarshal(bodyBytes, target)
}

func (c *client) getRequestHeader(reqOpts ...RequestOption) map[string]string {
	if len(reqOpts) == 0 {
		return nil
	}
	reqOpt := reqOpts[0]
	header := reqOpt.Header
	if header == nil {
		header = make(map[string]string)
	}
	if reqOpt.GroupPath != "" {
		header[GroupPathHeader] = reqOpt.GroupPath
	}
	return header
}

func getContextIDFromCtx(ctx context.Context) string {
	dsbContext := contxt.NewDSBContext(ctx)
	return dsbContext.GetRequestID()
}
