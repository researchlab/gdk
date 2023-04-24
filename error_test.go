package gdk

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestErrorCause(t *testing.T) {
	t.Run("语法检测", func(t *testing.T) {
		f := func() error {
			return ErrorCause(fmt.Errorf("error"))
		}
		err := f()
		if err == nil {
			t.Errorf("got nil, want error")
		}
		if err.Error() != "error" {
			t.Errorf("got %v, want error", err.Error())
		}
		err2 := ErrorCause(err)
		if err2 == nil {
			t.Errorf("got nil, want error")
		}
		if err2.Error() != "error" {
			t.Errorf("got %v, want error", err2.Error())
		}
		// 检测 err 是不是 error类型
		if _, ok := err.(error); !ok {
			t.Errorf("got false, want true, err")
		}
		// err 同时实现了 Err 接口， 所以应该也是Err 类型
		if _, ok := err.(Err); !ok {
			t.Errorf("got false, want true, Err")
		}
		// 检测err对象是否实现了WithTags(string)Err方法
		// 因为err实现了Err接口， 所以它应该实现了Err接口中的WithTags方法
		if _, ok := err.(interface {
			WithTag(string) Err
		}); !ok {
			t.Errorf("got false, want true, WithTag")
		}
	})
}

func foo(i int) error {
	if i == 0 {
		return ErrorCause(fmt.Errorf("testError"))
	}
	return foo(i - 1)
}

func TestIs(t *testing.T) {
	var (
		ERR_NUM_INVALID = 1000
		ERR_TEST_CODE   = 1001
	)
	err := ErrorCause(fmt.Errorf("error")).WithCode(ERR_NUM_INVALID)
	if !err.Is(ERR_NUM_INVALID) {
		t.Errorf("got %v, want error.code=1000", err.Detail())
	}
	if err.Is(ERR_TEST_CODE) {
		t.Errorf("got %v, want error.code !=1001", err.Detail())
	}
}

func TestWithTag(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		err := ErrorCause(fmt.Errorf("error")).WithTag("gdk").Export()
		if err.Tag != "gdk" {
			t.Errorf("got %v, want tag=gdk", err.Tag)
		}
	})

	t.Run("Global Tag", func(t *testing.T) {
		SetGlobalTag("globalGDK")
		err := ErrorCause(fmt.Errorf("error")).WithTag("gdk").Export()
		if err.Tag != "gdk" && err.GlobalTag != "globalGDK" {
			t.Errorf("got tag=%v, globalTag=%v, want tag=gdk, globalTag=globalGDK", err.Tag, err.GlobalTag)
		}
		err1 := ErrorCause(fmt.Errorf("error")).Export()
		if err1.Tag != "" && err1.GlobalTag != "globalGDK" {
			t.Errorf("got tag=%v, globalTag=%v, want tag=, globalTag=globalGDK", err.Tag, err.GlobalTag)
		}
	})
}

func TestWithFields(t *testing.T) {
	globalFields := map[string]interface{}{
		"service": "job",
	}
	fields := map[string]interface{}{
		"version": "1.0.0",
		"build":   20231101,
	}
	t.Run("Fields", func(t *testing.T) {
		err := ErrorCause(fmt.Errorf("error")).WithFields(fields).Export()
		if !reflect.DeepEqual(err.Fields, fields) {
			t.Errorf("got %v, want %v", err.Fields, fields)
		}
	})
	t.Run("GlobalFields", func(t *testing.T) {
		SetGlobalFields(globalFields)
		err := ErrorCause(fmt.Errorf("error")).WithFields(fields).Export()
		if !reflect.DeepEqual(err.Fields, fields) {
			t.Errorf("got %v, want %v", err.Fields, fields)
		}
		if !reflect.DeepEqual(err.GlobalFields, globalFields) {
			t.Errorf("got %v, want %v", err.GlobalFields, globalFields)
		}
	})
}

func TestWithCode(t *testing.T) {
	t.Run("code", func(t *testing.T) {
		ERR_INVALID := 1000
		err := ErrorCause(fmt.Errorf("error")).WithCode(ERR_INVALID)
		if !err.Is(ERR_INVALID) {
			t.Errorf("got %v, want code=1000", err.Detail())
		}
	})
}

func TestError(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		err := ErrorCause(fmt.Errorf("error"))
		if err.Error() != "error" {
			t.Errorf("got %v, want error", err.Error())
		}
	})
	t.Run("递归错误", func(t *testing.T) {
		err := foo(3)
		if err == nil {
			t.Errorf("got nil, want !=nil")
		}
		if err.Error() != "testError" {
			t.Errorf("got %v, want testError", err.Error())
		}
	})

}

func TestExport(t *testing.T) {
	t.Run("Export", func(t *testing.T) {
		code := 1000
		tag := "machine"
		fields := map[string]interface{}{
			"inputs":   []int{1, 1, 2},
			"outs":     nil,
			"location": "db connection refused",
		}
		err := ErrorCause(fmt.Errorf("error")).WithTag(tag).WithFields(fields).WithCode(code)
		e := err.Export()
		if e.Code != code {
			t.Errorf("got %v, want %v", e.Code, code)
		}
		if e.Tag != tag {
			t.Errorf("got %v, want %v", e.Tag, tag)
		}
		if !reflect.DeepEqual(e.Fields, fields) {
			t.Errorf("got %v, want %v", e.Fields, fields)
		}
	})
}

func TestDetail(t *testing.T) {
	t.Run("Detail", func(t *testing.T) {
		code := 1000
		tag := "machine"
		fields := map[string]interface{}{
			"inputs":   []interface{}{"a", "b"}, // 注意这里不能使用[]string{"a","b"} 否则下面json.Unmarshal反序列化之后时[]interface{} 此时因为类型不同，DeepEqual()=false
			"outs":     "1011",
			"location": "db connection refused",
		}
		err := ErrorCause(fmt.Errorf("error")).WithTag(tag).WithFields(fields).WithCode(code)

		d := err.Detail()
		var e ErrDetail
		ee := json.Unmarshal([]byte(d), &e)

		//decoder := json.NewDecoder(bytes.NewReader([]byte(d)))
		//// 未设置UseNumber, 长整型会丢失精度
		//decoder.UseNumber()
		//ee := decoder.Decode(&e)
		if ee != nil {
			t.Errorf("json.Unmarshal error %v", ee)
		}
		if c, ok := e.Code.(float64); !ok {
			t.Log("type:", reflect.TypeOf(e.Code))
			t.Errorf("got %v, want %v, ok %v", e.Code, code, ok)
		} else if int(c) != code {
			t.Errorf("got %v, want %v", e.Code, code)
		}

		if e.Tag != tag {
			t.Errorf("got %v, want %v", e.Tag, tag)
		}
		if !reflect.DeepEqual(e.Fields, fields) {
			t.Errorf("got %v, want %v", e.Fields, fields)
		}
		for k, v := range e.Fields {
			if !reflect.DeepEqual(v, fields[k]) {
				t.Errorf("key:%v, v:%v, want:%v, cmp:%v", k, v, fields[k], v == fields[k])
			}
		}
	})
}

func TestErrorf(t *testing.T) {
	t.Run("Errorf", func(t *testing.T) {
		const ERR_INVALID = 1000
		err := Errorf("error").WithCode(ERR_INVALID)
		if !err.Is(ERR_INVALID) {
			t.Errorf("got %v, want 1000", err.Detail())
		}
		if err.Error() != "error" {
			t.Errorf("got %v, want error", err.Error())
		}
	})
}

func TestErrorT(t *testing.T) {
	t.Run("ErrorT with code[int]", func(t *testing.T) {
		const (
			ERR_UNMARSHAL_INVALID = 1000
			ERR_PARAMS_INVALID    = 1001
		)
		errorTemplates := map[any]string{
			ERR_UNMARSHAL_INVALID: "json.Unmarshal error %+v, inputs:%+v",
			ERR_PARAMS_INVALID:    "%v invalid",
		}
		SetGlobalErrorTemplates(errorTemplates)
		err := ErrorT(ERR_PARAMS_INVALID, "stu.address")
		if !err.Is(ERR_PARAMS_INVALID) {
			t.Errorf("got %v, want %v", err.Detail(), ERR_PARAMS_INVALID)
		}
		if err.Error() != "stu.address invalid" {
			t.Errorf("got %v, want %v", err.Error(), "stu.address invalid")
		}
		err2 := ErrorT(ERR_UNMARSHAL_INVALID, fmt.Errorf("cannot covert object into golang struct"), "{\"address\":\"china\",\"code\":1000\"}")
		if !err2.Is(ERR_UNMARSHAL_INVALID) {
			t.Errorf("got %v, want %v", err.Detail(), ERR_UNMARSHAL_INVALID)
		}
		want := "json.Unmarshal error cannot covert object into golang struct, inputs:{\"address\":\"china\",\"code\":1000\"}"
		if err2.Error() != want {
			t.Errorf("got %v, want %v", err2.Error(), want)
		}
	})
	t.Run("ErrorT with code[string]", func(t *testing.T) {
		const (
			ERR_UNMARSHAL_INVALID = "ERR_UNMARSHAL_INVALID"
			ERR_PARAMS_INVALID    = "ERR_PARAMS_INVALID"
		)
		errorTemplates := map[any]string{
			ERR_UNMARSHAL_INVALID: "json.Unmarshal error %+v, inputs:%+v",
			ERR_PARAMS_INVALID:    "%v invalid",
		}
		SetGlobalErrorTemplates(errorTemplates)
		err := ErrorT(ERR_PARAMS_INVALID, "stu.address")
		if !err.Is(ERR_PARAMS_INVALID) {
			t.Errorf("got %v, want %v", err.Detail(), ERR_PARAMS_INVALID)
		}
		if err.Error() != "stu.address invalid" {
			t.Errorf("got %v, want %v", err.Error(), "stu.address invalid")
		}
		err2 := ErrorT(ERR_UNMARSHAL_INVALID, fmt.Errorf("cannot covert object into golang struct"), "{\"address\":\"china\",\"code\":1000\"}")
		if !err2.Is(ERR_UNMARSHAL_INVALID) {
			t.Errorf("got %v, want %v", err.Detail(), ERR_UNMARSHAL_INVALID)
		}
		want := "json.Unmarshal error cannot covert object into golang struct, inputs:{\"address\":\"china\",\"code\":1000\"}"
		if err2.Error() != want {
			t.Errorf("got %v, want %v", err2.Error(), want)
		}
	})
}

func TestErrorText(t *testing.T) {
	const ERR_PARSE_FAILED = "PARSE_FAILED"
	_, e := json.Marshal(func() {})
	SetGlobalTag("ip:192.168.1.12")
	SetGlobalFields(map[string]interface{}{"build": 20231121})
	err := ErrorCause(e).WithCode(ERR_PARSE_FAILED).WithTag("FunctionMarshal").WithFields(map[string]interface{}{"key": "value"})
	want := `CallChains=TestErrorText, Tag=FunctionMarshal, GlobalFields={"build":20231121}, Fields={"key":"value"}, Code=PARSE_FAILED, Error=json: unsupported type: func() `
	if err.ErrorText() != want {
		t.Errorf("got %v, want %v", err.ErrorText(), want)
	}
}
