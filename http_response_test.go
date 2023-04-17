package gdk

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
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
		name    string
		args    args
		wantR   *Response
		wantErr bool
	}{
		{name: "nil response", args: args{resp: nil}, wantR: &Response{}, wantErr: true},
		{name: "json response 200", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{"Content-Type": "application/json"},
			jsonb,
		)}, wantR: &Response{body: []byte("{\"name\":\"mike\"}")}, wantErr: false},
		{name: "gzip response 200", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{
				"Content-Type":     "application/json",
				"Content-Encoding": "gzip",
			},
			gzipb.Bytes(),
		)}, wantR: &Response{body: []byte("{\"name\":\"mike\"}")}, wantErr: false},
		{name: "gzip NewReader error", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{
				"Content-Type":     "application/json",
				"Content-Encoding": "gzip",
			},
			//gzipb.Bytes(),
			jsonb,
		)}, wantR: &Response{}, wantErr: true},
		{name: "ioutil ReadAll error", args: args{resp: httpResponse(
			http.StatusOK,
			map[string]string{
				"Content-Type":     "application/json",
				"Content-Encoding": "gzip",
			},
			gzipbad,
		)}, wantR: &Response{body: []byte("{\"name\":\"mike\"}")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := ResponseRecorder(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResponseRecorder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR.String() != tt.wantR.String() {
				t.Errorf("ResponseRecorder() = %v, want %v", gotR, tt.wantR)
			}

		})
	}
}

func TestResponseBody(t *testing.T) {
	type fields struct {
		body []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{name: "nil body", fields: fields{body: nil}, want: []byte{}},
		{name: "not nil body", fields: fields{body: []byte("mike")}, want: []byte("mike")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				body: tt.fields.body,
			}
			if got := r.Body(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Body() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseStatus(t *testing.T) {
	type fields struct {
		status string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "OK status", fields: fields{status: "OK"}, want: "OK"},
		{name: "OK 200 status", fields: fields{status: "OK 200"}, want: "OK 200"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				status: tt.fields.status,
			}
			if got := r.Status(); got != tt.want {
				t.Errorf("Response.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseStatusCode(t *testing.T) {
	type fields struct {
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "status code 200", fields: fields{statusCode: 200}, want: 200},
		{name: "status code 500", fields: fields{statusCode: 500}, want: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				statusCode: tt.fields.statusCode,
			}
			if got := r.StatusCode(); got != tt.want {
				t.Errorf("Response.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseHeader(t *testing.T) {
	type fields struct {
		header http.Header
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{name: "nil header", fields: fields{header: nil}, want: http.Header{}},
		{name: "positive header", fields: fields{header: http.Header(map[string][]string{"name": []string{"mike"}})}, want: map[string][]string{"name": []string{"mike"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				header: tt.fields.header,
			}
			if got := r.Header(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseString(t *testing.T) {
	type fields struct {
		body []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "nil body", fields: fields{body: nil}, want: ""},
		{name: "positive body", fields: fields{body: []byte("body")}, want: "body"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				body: tt.fields.body,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Response.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseSize(t *testing.T) {
	type fields struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "response size 100", fields: fields{size: 100}, want: 100},
		{name: "response size 0", fields: fields{size: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				size: tt.fields.size,
			}
			if got := r.Size(); got != tt.want {
				t.Errorf("Response.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseUnmarshal(t *testing.T) {
	type fields struct {
		body []byte
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "body not in json format", fields: fields{body: []byte("{11,12}")}, args: args{data: new(interface{})}, wantErr: true},
		{name: "body in json format", fields: fields{body: []byte("{\"name\":\"mike\"}")}, args: args{data: new(interface{})}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				body: tt.fields.body,
			}
			if err := r.Unmarshal(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Response.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
