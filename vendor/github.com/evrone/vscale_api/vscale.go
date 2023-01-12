package vscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	defaultApiEndpoint = "https://api.vscale.io"
	mediaType          = "application/json"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	token   string

	// Vscale services
	Account        AccountService
	Background     BackgroundService
	Configurations ConfigurationsService
	Notifications  NotificationsService
	Scalet         ScaletService
	SSHKey         SSHService

	// Optional
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response *http.Response
	Message  string
}

type ArgError struct {
	arg, reason string
}

var _ error = &ArgError{}

func NewArgError(arg, reason string) *ArgError {
	return &ArgError{
		arg:    arg,
		reason: reason,
	}
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%s is invalid because %s", e.arg, e.reason)
}

func New(token string) *Client {
	return NewClient(http.DefaultClient, token)
}

func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(defaultApiEndpoint)

	c := &Client{client: httpClient, BaseURL: baseUrl, token: token}

	c.Account = &AccountServiceOp{client: c}
	c.Background = &BackgroundServiceOp{client: c}
	c.Configurations = &ConfigurationsServiceOp{client: c}
	c.Notifications = &NotificationsServiceOp{client: c}
	c.Scalet = &ScaletServiceOp{client: c}
	c.SSHKey = &SSHServiceOp{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	buf := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("X-Token", c.token)

	return req, nil
}

func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	return &response
}

func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err != nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	if err = CheckResponse(resp); err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			if _, err := io.Copy(w, resp.Body); err != nil {
				return nil, err
			}
		} else {
			if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()
	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}
	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}
