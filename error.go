package gdk

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type Err interface {
	Is(int) bool
	WithTag(string) Err
	WithFields(map[string]interface{}) Err
	WithCode(int) Err
	Error() string
	Export() ErrDetail
	Detail() string
}

var (
	globalTag    string
	globalFields map[string]interface{}
)

func init() {
	globalFields = make(map[string]interface{})
}

// SetGlobalTag global tag
func SetGlobalTag(globalTag string) {
	globalTag = globalTag
}

// SetGlobalFields global fields
func SetGlobalFields(fields map[string]interface{}) {
	for k, v := range fields {
		globalFields[k] = v
	}
}

// err error struct
type err struct {
	chains []string
	tag    string
	code   int
	fields map[string]interface{}
	e      error
}

// ErrDetail error detail struct
type ErrDetail struct {
	Chains       []string               `json:"CallChains,omitempty"` // 反序列化时,如果该字段为空,则不进行序列化输出
	GlobalTag    string                 `json:"GlobalTag,omitempty"`
	Tag          string                 `json:"Tag,omitempty"`
	GlobalFields map[string]interface{} `json:"GlobalFields,omitempty"`
	Fields       map[string]interface{} `json:"Fields,omitempty"`
	Code         int                    `json:"Code,omitempty"`
	E            string                 `json:"Error,omitempty"`
	e            error                  `json:"-"`
}

// Error return error string
func (e *err) Error() string {
	if e.e == nil {
		return ""
	}
	return e.e.Error()
}

// Export  export error detail
func (e *err) Export() ErrDetail {
	return ErrDetail{
		Chains:       e.chains,
		GlobalTag:    globalTag,
		Tag:          e.tag,
		GlobalFields: globalFields,
		Fields:       e.fields,
		Code:         e.code,
		E: func() string {
			if e.e == nil {
				return ""
			}
			return e.e.Error()
		}(),
	}
}

// Detail details of error by json format
func (e *err) Detail() string {
	b, _ := json.Marshal(e.Export())
	return string(b)
}

// Is compare two error code , return true if equals
func (e *err) Is(code int) bool {
	return e.code == code
}

// WithTag  set tag for the given error
func (e *err) WithTag(tag string) Err {
	e.tag = tag
	return e
}

// WithFields with more error messages for the given error
func (e *err) WithFields(fields map[string]interface{}) Err {
	for k, v := range fields {
		e.fields[k] = v
	}
	return e
}

// WithCode  error code
func (e *err) WithCode(code int) Err {
	e.code = code
	return e
}

// ErrorCause error recorder
func ErrorCause(e error) Err {
	v, ok := e.(*err)
	if ok {
		v.chains = append([]string{callerName()}, v.chains...)
		return v
	}
	return &err{
		chains: []string{callerName()},
		e:      e,
		fields: make(map[string]interface{}),
	}
}

func Errorf(format string, a ...any) Err {
	return &err{
		chains: []string{callerName()},
		e:      fmt.Errorf(format, a...),
		fields: make(map[string]interface{}),
	}
}

// callerName return parent function name, if not exits return #
func callerName() (caller string) {
	pc, _, _, ok := runtime.Caller(2) // 0: function-self, 1: parent function caller
	if !ok {
		caller = "#"
	} else {
		path := runtime.FuncForPC(pc).Name()
		items := strings.Split(path, ".")
		caller = items[len(items)-1]
		if len(caller) == 0 {
			caller = path
		}
	}
	return caller
}
