package gdk

import (
	"compress/gzip"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")

// ReadResponse return http.Response.body, isGzip=true if contentEncoding=gzip
func ReadResponse(resp *http.Response) (bodyBytes []byte, isGzip bool, err error) {
	if resp == nil {
		err = errors.New("nil http response")
		return
	}
	body := resp.Body
	// check gzip
	if strings.EqualFold(resp.Header.Get(hdrContentEncodingKey), "gzip") && resp.ContentLength != 0 {
		isGzip = true
		if _, ok := body.(*gzip.Reader); !ok {
			body, err = gzip.NewReader(body)
			if err != nil {
				return
			}
		}
	}
	defer body.Close()
	if bodyBytes, err = ioutil.ReadAll(body); err != nil {
		return
	}
	return bodyBytes, isGzip, err
}
