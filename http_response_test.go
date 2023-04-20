package gdk

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

// named-return, to avoid side effects of defer opt
func gzipFast(a []byte) (b bytes.Buffer) {

	gz := gzip.NewWriter(&b)
	defer gz.Close()
	if _, err := gz.Write(a); err != nil {
		panic(err)
	}
	gz.Flush()
	return b
}

// badgzip
func gzipBad(a []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	defer gz.Close()
	if _, err := gz.Write(a); err != nil {
		panic(err)
	}
	gz.Flush()
	return b.Bytes()
}

func httpResponse(statusCode int, headers map[string]string, body []byte) *http.Response {
	header := http.Header{}
	for name, value := range headers {
		header.Set(name, value)
	}
	return &http.Response{
		Status:           strconv.Itoa(statusCode) + " " + http.StatusText(statusCode),
		StatusCode:       statusCode,
		Proto:            "HTTP/1.0",
		ProtoMajor:       1,
		ProtoMinor:       0,
		Header:           header,
		Body:             ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength:    int64(len(body)),
		TransferEncoding: []string{},
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		TLS:              nil,
	}
}

func TestResponseRecorder(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	jsonb, _ := json.Marshal(map[string]interface{}{"name": "mike"})
	gzipb := gzipFast(jsonb)
	gzipbad := gzipBad(jsonb)

	tests := []struct {
		name       string
		args       args
		wantBody   []byte
		wantIsGzip bool
		wantErr    bool
	}{
		{name: "nil response", args: args{resp: nil}, wantBody: []byte{}, wantIsGzip: false, wantErr: true},
		{name: "json response 200", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{"Content-Type": "application/json"},
			jsonb,
		)}, wantBody: []byte("{\"name\":\"mike\"}"), wantIsGzip: false, wantErr: false},
		{name: "gzip response 200", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{
				"Content-Type":     "application/json",
				"Content-Encoding": "gzip",
			},
			gzipb.Bytes(),
		)}, wantBody: []byte("{\"name\":\"mike\"}"), wantIsGzip: true, wantErr: false},
		{name: "gzip NewReader error", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{
				"Content-Type":     "application/json",
				"Content-Encoding": "gzip",
			},
			//gzipb.Bytes(),
			jsonb,
		)}, wantBody: []byte{}, wantIsGzip: true, wantErr: true},
		{name: "ioutil ReadAll error", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{
				"Content-Type":     "application/json",
				"Content-Encoding": "gzip",
			},
			gzipbad,
		)}, wantBody: []byte("{\"name\":\"mike\"}"), wantIsGzip: true, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyBytes, isGzip, err := ReadResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(bodyBytes) != string(tt.wantBody) {
				t.Errorf("ReadResponse() = %s, want %s", bodyBytes, tt.wantBody)
			}
			if isGzip != tt.wantIsGzip {
				t.Errorf("ReadResponse()  isGizp=%v, want %v", isGzip, tt.wantIsGzip)
			}
		})
	}
}
