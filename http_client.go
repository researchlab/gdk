package gdk

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	FORM_MULTIPART   = "multipart/form-data"
	FORM_ENCODED     = "application/x-www-form-urlencoded"
	APPLICATION_JSON = "application/json"
	CONTENT_TYPE     = "Content-Type"
)

// FileUploadInfo upload file info struct
type FileUploadInfo struct {
	Name     string // form name
	Filepath string
	FileName string
}

type HttpOptions struct {
	C           *http.Client
	Url         string
	Params      map[string]string
	Headers     map[string]string
	Files       []FileUploadInfo // special for HttpPostFile
	contentType string
}

const (
	ERR_PARAMS_INVALID      = 11111
	ERR_OPEN_FILE_FAILED    = 11112
	ERR_WRITER_FAILED       = 11113
	ERR_WRITE_FIELD_FAILED  = 11114
	ERR_WRITER_CLOSE_FAILED = 11115
	ERR_UNKOWN_TYPE         = 11116
)

// HttpGet request to target url
func HttpGet(ho *HttpOptions) (*http.Response, error) {
	if ho == nil {
		return nil, Errorf("Params invalid").WithCode(ERR_PARAMS_INVALID)
	}
	if ho.C == nil {
		ho.C = http.DefaultClient
	}
	urlParams := url.Values{}
	Url, _ := url.Parse(ho.Url)
	for k, v := range ho.Params {
		urlParams.Set(k, v)
	}

	Url.RawQuery = urlParams.Encode()
	urlPath := Url.String()

	httpReq, _ := http.NewRequest(http.MethodGet, urlPath, nil)

	for k, v := range ho.Headers {
		httpReq.Header.Add(k, v)
	}
	return ho.C.Do(httpReq)
}

// HttpPostJSON   post json format value to http server
func HttpPostJSON(ho *HttpOptions) (*http.Response, error) {
	ho.contentType = APPLICATION_JSON
	return doPost(ho)
}

// HttpPostForm post form field value
func HttpPostForm(ho *HttpOptions) (*http.Response, error) {
	ho.contentType = FORM_ENCODED
	return doPost(ho)
}

// HttpPostFiles post files
func HttpPostFiles(ho *HttpOptions) (*http.Response, error) {
	ho.contentType = FORM_MULTIPART
	return doPost(ho)
}

func doPost(ho *HttpOptions) (*http.Response, error) {
	if ho.C == nil {
		ho.C = http.DefaultClient
	}
	body, realContentType, err := buildRequest(ho)
	if err != nil {
		return nil, ErrorCause(err)
	}
	req, _ := http.NewRequest(http.MethodPost, ho.Url, body)
	req.Header.Add(CONTENT_TYPE, realContentType)
	for k, v := range ho.Headers {
		req.Header.Add(k, v)
	}
	return ho.C.Do(req)
}

func buildRequest(ho *HttpOptions) (io.Reader, string, error) {
	switch ho.contentType {
	case APPLICATION_JSON:
		bytesData, _ := json.Marshal(ho.Params)
		return bytes.NewReader(bytesData), ho.contentType, nil
	case FORM_ENCODED:
		urlValues := url.Values{}
		for k, v := range ho.Params {
			urlValues.Set(k, v)
		}
		body := urlValues.Encode()
		return strings.NewReader(body), ho.contentType, nil
	case FORM_MULTIPART:
		if ho.Files == nil || len(ho.Files) == 0 {
			return nil, ho.contentType, Errorf("Files invalid").WithCode(ERR_PARAMS_INVALID)
		}
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		for _, uploadFile := range ho.Files {
			file, err := os.Open(uploadFile.Filepath)
			if err != nil {
				return nil, "", ErrorCause(err).WithCode(ERR_OPEN_FILE_FAILED)
			}
			fileName := filepath.Base(uploadFile.Filepath)
			if uploadFile.FileName != "" {
				fileName = uploadFile.FileName
			}
			part, err := writer.CreateFormFile(uploadFile.Name, fileName)
			if err != nil {
				return nil, "", ErrorCause(err).WithCode(ERR_WRITER_FAILED)
			}
			io.Copy(part, file)
			file.Close()
		}
		for k, v := range ho.Params {
			if err := writer.WriteField(k, v); err != nil {
				return nil, "", ErrorCause(err).WithCode(ERR_WRITE_FIELD_FAILED)
			}
		}
		if err := writer.Close(); err != nil {
			return nil, "", ErrorCause(err).WithCode(ERR_WRITER_CLOSE_FAILED)
		}
		return body, writer.FormDataContentType(), nil

	default:
		return nil, "unkownType", Errorf("unkownType").WithCode(ERR_UNKOWN_TYPE)
	}
}
