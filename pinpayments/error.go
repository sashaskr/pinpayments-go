package pinpayments

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Content  string         `json:"content,omitempty"`
	Response *http.Response `json:"response"` // the full response that produced the error
}

var (
	errEmptyApiKey         = errors.New("you must provide a non-empty API key")
	errEmptyApiPublishable = errors.New("you mush provide a non-empty API Password")
	errBadBaseURL          = errors.New("malformed base url, it must contain a trailing slash")
)

func (e *Error) Error() string {
	return fmt.Sprintf("response failed with status %s\npayload: %v", e.Message, e.Content)
}

func newError(r *http.Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		e.Content = string(c)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	return &e
}
