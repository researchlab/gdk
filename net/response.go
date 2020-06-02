package net

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")

// Response is an object recorder executed request and its values.
type Response struct {
	status     string // e.g. "200 OK"
	statusCode int    // e.g. 200
	//Proto      string // e.g. "HTTP/1.0"
	//ProtoMajor int    // e.g. 1
	//ProtoMinor int    // e.g. 0

	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	header http.Header
	body   []byte
	size   int64
}

// ResponseRecorder   is an object recorder the given http response
func ResponseRecorder(resp *http.Response) (r *Response, err error) {
	r = &Response{}
	body := resp.Body

	// check gzip
	if strings.EqualFold(resp.Header.Get(hdrContentEncodingKey), "gzip") && resp.ContentLength != 0 {
		if _, ok := body.(*gzip.Reader); !ok {
			body, err = gzip.NewReader(body)
			if err != nil {
				return
			}
			defer body.Close()
		}
	}
	// copy body
	if r.body, err = ioutil.ReadAll(body); err != nil {
		return
	}

	// copy status
	r.status = resp.Status
	r.statusCode = resp.StatusCode

	// copy Header
	headers := make(map[string][]string)
	for name, values := range resp.Header {
		headers[name] = values
	}
	r.header = http.Header(headers)
	r.size = int64(len(r.body))
	return
}

// Body method returns HTTP response as []byte array for the executed request.
func (r *Response) Body() []byte {
	if r.body == nil {
		return []byte{}
	}
	return r.body
}

// Status method returns the HTTP status string for the executed request.
// Example: 200 OK
func (r *Response) Status() string {
	return r.status
}

// StatusCode method returns the HTTP status code for the executed request.
// Example: 200
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Header method returns the response headers
func (r *Response) Header() http.Header {
	if r.header == nil {
		return http.Header{}
	}
	return r.header
}

// String method returns the body of the server response as String.
func (r *Response) String() string {
	if r.body == nil {
		return ""
	}
	return strings.TrimSpace(string(r.body))
}

// Size method returns the HTTP response size in bytes. Ya, you can relay on HTTP `Content-Length` header,
// however it won't be good for chucked transfer/compressed response. Since ResponseRecorder calculates response size
// at the client end. You will get actual size of the http response.
func (r *Response) Size() int64 {
	return r.size
}

// Unmarshal json decode response to data
func (r *Response) Unmarshal(data interface{}) error {
	return json.Unmarshal(r.body, data)
}
