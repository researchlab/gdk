package gdk

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpGet(t *testing.T) {
	var want = "Hello, client"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, want) }))
	defer ts.Close()
	res, err := HttpGet(&HttpOptions{Url: ts.URL})
	if err != nil {
		t.Errorf("HttpGet %+v", err)
		return
	}
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ReadBody %v", err)
		return
	}
	res.Body.Close()
	if string(got) != want {
		t.Errorf("got %s, want %v", got, want)
	}
}
