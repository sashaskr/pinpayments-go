package pinpayments

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	*http.Response
	content []byte
}

func (c *Client) NewAPIRequest(secret bool, method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	if c.config.testing {
		u.Query().Add("testmode", "true")
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)

	if c.page != nil {
		q := req.URL.Query()
		q.Add("page", strconv.Itoa(*c.page))
		req.URL.RawQuery = q.Encode()
	}

	if err != nil {
		return nil, err
	}

	var token string

	if secret {
		token = base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{c.secretKey, ""}, ":")))
	} else {
		token = base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{c.publishableKey, ""}, ":")))
	}

	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, token}, " "))
	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestAccept)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Connection", Connection)
	return
}

func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	return response, nil
}

func newResponse(r *http.Response) (*Response, error) {
	var res Response
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		res.content = c
	}
	err = json.NewDecoder(r.Body).Decode(&res)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	res.Response = r
	return &res, err
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}
