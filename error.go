package gdk

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type Err interface {
	Is(any) bool
	WithTag(string) Err
	WithFields(map[string]interface{}) Err
	WithCode(any) Err
	Error() string
	Export() ErrDetail
	Detail() string
	DetailText() string
}

var (
	globalTag            string
	globalFields         map[string]interface{}
	globalErrorTemplates map[any]string
)

const (
	GTAG            = "globalTag"
	GFIELDS         = "globalFields"
	GERRORTEMPLATES = "globalErrorTemplates"
)

func init() {
	globalFields = make(map[string]interface{})
	globalErrorTemplates = make(map[any]string)
}

func UnsetGlobals(options ...string) {
	if len(options) != 0 {
		for _, v := range options {
			switch v {
			case GTAG:
				globalTag = ""
			case GFIELDS:
				globalFields = make(map[string]interface{})
			case GERRORTEMPLATES:
				globalErrorTemplates = make(map[any]string)
			default:
			}
		}
		return
	}
	globalFields = make(map[string]interface{})
	globalErrorTemplates = make(map[any]string)
	globalTag = ""
}

// SetGlobalErrorTemplates cache error templates
// 建议做多设置一次
func SetGlobalErrorTemplates(templates map[any]string) {
	for k, v := range templates {
		globalErrorTemplates[k] = v
	}
}

// SetGlobalTag global tag
func SetGlobalTag(tag string) {
	globalTag = tag
}

// SetGlobalFields global fields, set at most once
// 建议最多设置一次
func SetGlobalFields(fields map[string]interface{}) {
	for k, v := range fields {
		globalFields[k] = v
	}
}

// err error struct
type err struct {
	chains []string
	tag    string
	code   any
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
	Code         any                    `json:"Code,omitempty"`
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

// DetailText textplain of the error
func (e *err) DetailText() string {
	var text string
	if e.chains != nil && len(e.chains) != 0 {
		text = fmt.Sprintf("CallChains=%s, ", strings.Join(e.chains, "."))
	}
	if len(globalTag) != 0 {
		text += fmt.Sprintf("GlobalTag=%s, ", globalTag)
	}
	if len(e.tag) != 0 {
		text += fmt.Sprintf("Tag=%s, ", e.tag)
	}
	if globalFields != nil && len(globalFields) != 0 {
		b, _ := json.Marshal(globalFields)
		text += fmt.Sprintf("GlobalFields=%s, ", b)
	}
	if e.fields != nil && len(e.fields) != 0 {
		b, _ := json.Marshal(e.fields)
		text += fmt.Sprintf("Fields=%s, ", b)
	}
	text += fmt.Sprintf("Code=%+v, ", e.code)
	if e.e != nil {
		text += fmt.Sprintf("Error=%s ", e.e.Error())
	}
	return text
}

// Is compare two error code , return true if equals
func (e *err) Is(code any) bool {
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
func (e *err) WithCode(code any) Err {
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

// Errorf new error with format
func Errorf(format string, a ...any) Err {
	return &err{
		chains: []string{callerName()},
		e:      fmt.Errorf(format, a...),
		fields: make(map[string]interface{}),
	}
}

// ErrorT new error by error code and error template
func ErrorT(code any, a ...any) Err {
	format, ok := globalErrorTemplates[code]
	var e2 error
	if ok {
		e2 = fmt.Errorf(format, a...)
	} else {
		emsg := ""
		for _, v := range a {
			emsg += fmt.Sprintf("%+v ", v)
		}
		e2 = fmt.Errorf("%s", emsg)
	}
	return &err{
		chains: []string{callerName()},
		e:      e2,
		code:   code,
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
